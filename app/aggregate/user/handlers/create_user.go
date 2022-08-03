package handlers

import (
	"fmt"
	"net/http"

	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/common"
	"github.com/go-chi/render"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var (
	ErrInvalidName = common.NewApplicationError(
		http.StatusBadRequest,
		"Name cannot be empty",
	)
	ErrInvalidUsername = common.NewApplicationError(
		http.StatusBadRequest,
		"Username cannot be empty",
	)
)

func (r CreateUserRequest) Bind(_ *http.Request) error {
	if r.Name == "" {
		return ErrInvalidName
	}

	if r.Username == "" {
		return ErrInvalidUsername
	}

	if _, err := userEntity.ParseEmail(r.Email); err != nil {
		return err
	}
	return nil
}

func (h *UserHTTPHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	request := &CreateUserRequest{}

	if err := render.Bind(r, request); err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	id, err := h.service.CreateUser(request.Name, request.Username, request.Email)
	if err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	render.Status(r, http.StatusCreated)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)
}
