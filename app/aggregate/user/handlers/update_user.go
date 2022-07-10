package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

var ErrEmptyUpdateUserRequest = errors.New("you must pass at least one field")

func (r UpdateUserRequest) Bind(_ *http.Request) error {
	if r.Name == nil && r.Email == nil {
		return ErrEmptyUpdateUserRequest
	}

	if r.Email != nil {
		if _, err := mail.ParseAddress(*r.Email); err != nil {
			return ErrInvalidEmail
		}
	}
	return nil
}

func (h *UserHTTPHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	request := &UpdateUserRequest{}

	id, err := getUserIDFromURL(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch err := render.Bind(r, request); err {
	case nil:
	case ErrEmptyUpdateUserRequest, ErrInvalidEmail:
		w.WriteHeader(http.StatusBadRequest)
		return
	default:
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.service.UpdateUser.Handle(service.UpdateUser{
		ID:    id,
		Name:  request.Name,
		Email: request.Email,
	}); err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)
}
