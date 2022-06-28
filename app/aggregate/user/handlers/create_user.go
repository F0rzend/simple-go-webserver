package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	userService "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var ErrInvalidEmail = errors.New("invalid email")

func (r CreateUserRequest) Bind(_ *http.Request) error {
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return ErrInvalidEmail
	}
	return nil
}

func (h *UserHTTPHandlers) createUser(w http.ResponseWriter, r *http.Request) {
	request := &CreateUserRequest{}

	if err := render.Bind(r, request); err != nil {
		switch err {
		case ErrInvalidEmail:
			w.WriteHeader(http.StatusBadRequest)
		default:
			log.Error().Err(err).Send()
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	id, err := h.service.CreateUser.Handle(userService.CreateUser{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
	})
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)
}
