package types

import (
	"errors"
	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
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

type SetBTCPriceRequest struct {
	Price float64 `json:"price"`
}

func (r SetBTCPriceRequest) Bind(_ *http.Request) error {
	if r.Price <= 0 {
		return ErrInvalidPrice{Price: r.Price}
	}
	return nil
}

var (
	ErrInvalidAction = errors.New("invalid action")
	ErrInvalidAmount = errors.New("invalid amount")
)

type ChangeUSDBalanceRequest struct {
	Action string  `json:"action"`
	Amount float64 `json:"amount"`
}

func (r ChangeUSDBalanceRequest) Bind(_ *http.Request) error {
	if _, err := domain.NewUSDAction(r.Action); err != nil {
		return ErrInvalidAction
	}

	if r.Amount < 0 {
		return ErrInvalidAmount
	}

	return nil
}

type ChangeBTCBalanceRequest struct {
	Action string  `json:"action"`
	Amount float64 `json:"amount"`
}

func (r ChangeBTCBalanceRequest) Bind(_ *http.Request) error {
	if _, err := domain.NewBTCAction(r.Action); err != nil {
		return ErrInvalidAction
	}

	if r.Amount < 0 {
		return ErrInvalidAmount
	}

	return nil
}
