package types

import (
	"net/http"
	"net/mail"

	"github.com/go-chi/render"
)

var (
	_ render.Binder = CreateUserRequest{}
)

type CreateUserRequest struct {
	Name     string
	Username string
	Email    string
}

func (r CreateUserRequest) Bind(_ *http.Request) error {
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return ErrInvalidEmail{Email: r.Email}
	}
	return nil
}

type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

func (r UpdateUserRequest) Bind(_ *http.Request) error {
	if r.Email != nil {
		if _, err := mail.ParseAddress(*r.Email); err != nil {
			return ErrInvalidEmail{Email: *r.Email}
		}
	}
	return nil
}
