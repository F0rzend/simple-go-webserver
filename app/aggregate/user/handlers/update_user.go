package handlers

import (
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

var ErrEmptyUpdateUserRequest = common.NewApplicationError(
	http.StatusNotModified,
	"At least one field must be updated",
)

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

	id, err := h.getUserIDFromRequest(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := render.Bind(r, request); err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	if err := h.service.UpdateUser(service.UpdateUserCommand{
		UserID: id,
		Name:   request.Name,
		Email:  request.Email,
	}); err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)
}
