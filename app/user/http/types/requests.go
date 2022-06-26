package types

import (
	"errors"
	"net/http"
	"net/mail"

	"github.com/F0rzend/simple-go-webserver/app/bitcoin/domain"

	"github.com/go-chi/render"
)

var (
	_ render.Binder = CreateUserRequest{}

	ErrBadRequest = errors.New("validation error")
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r CreateUserRequest) Bind(_ *http.Request) error {
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return ErrBadRequest
	}
	return nil
}

type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

func (r UpdateUserRequest) Bind(_ *http.Request) error {
	if r.Name == nil && r.Email == nil {
		return ErrBadRequest
	}

	if r.Email != nil {
		if _, err := mail.ParseAddress(*r.Email); err != nil {
			return ErrBadRequest
		}
	}
	return nil
}

type ChangeUSDBalanceRequest struct {
	Action string  `json:"action"`
	Amount float64 `json:"amount"`
}

func (r ChangeUSDBalanceRequest) Bind(_ *http.Request) error {
	if r.Action == "" || r.Amount == 0 {
		return ErrBadRequest
	}

	if _, err := domain.NewUSDAction(r.Action); err != nil {
		return ErrBadRequest
	}

	return nil
}

type ChangeBTCBalanceRequest struct {
	Action string  `json:"action"`
	Amount float64 `json:"amount"`
}

func (r ChangeBTCBalanceRequest) Bind(_ *http.Request) error {
	if _, err := domain.NewBTCAction(r.Action); err != nil {
		return ErrBadRequest
	}

	return nil
}
