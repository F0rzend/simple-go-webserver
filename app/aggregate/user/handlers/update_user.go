package userhandlers

import (
	"fmt"
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/common"
	"github.com/go-chi/render"
)

type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

func (r UpdateUserRequest) Bind(_ *http.Request) error {
	if r.Name == nil && r.Email == nil {
		return common.NewValidationError("nothing to update, please provide name or email")
	}

	if r.Email != nil {
		if _, err := userentity.ParseEmail(*r.Email); err != nil {
			return common.NewValidationError("invalid email")
		}
	}
	return nil
}

func (h *UserHTTPHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	request := &UpdateUserRequest{}

	id, err := h.getUserIDFromRequest(r)
	if err != nil {
		return fmt.Errorf("failed to get user id from request: %w", err)
	}

	if err := render.Bind(r, request); err != nil {
		return fmt.Errorf("failed to bind request: %w", err)
	}

	err = h.service.UpdateUser(id, request.Name, request.Email)
	if common.IsFlaggedError(err, common.FlagInvalidArgument) {
		return common.NewValidationError(err.Error())
	}
	if common.IsFlaggedError(err, common.FlagNotFound) {
		return common.NewNotFoundError(fmt.Sprintf("user with id %d not found", id))
	}
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)

	return nil
}
