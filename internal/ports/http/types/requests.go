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

type ErrInvalidEmail struct {
	Email string
}

func (e ErrInvalidEmail) Error() string {
	return "invalid email: " + e.Email
}

func (r CreateUserRequest) Bind(_ *http.Request) error {
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return ErrInvalidEmail{Email: r.Email}
	}
	return nil
}
