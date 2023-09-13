package userhandlers

import (
	"fmt"
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/common"
	"github.com/go-chi/render"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r CreateUserRequest) Bind(_ *http.Request) error {
	if r.Name == "" {
		return common.NewValidationError("Name cannot be empty")
	}

	if r.Username == "" {
		return common.NewValidationError("Username cannot be empty")
	}

	_, err := userentity.ParseEmail(r.Email)
	if common.IsFlaggedError(err, common.FlagInvalidArgument) {
		return common.NewValidationError(err.Error())
	}
	if err != nil {
		return fmt.Errorf("error parsing email: %w", err)
	}

	return nil
}

func (h *UserHTTPHandlers) CreateUser(w http.ResponseWriter, r *http.Request) error {
	request := &CreateUserRequest{}

	if err := render.Bind(r, request); err != nil {
		return fmt.Errorf("error binding request: %w", err)
	}

	id, err := h.service.CreateUser(request.Name, request.Username, request.Email)
	if common.IsFlaggedError(err, common.FlagInvalidArgument) {
		return common.NewValidationError(err.Error())
	}
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	render.Status(r, http.StatusCreated)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)

	return nil
}
