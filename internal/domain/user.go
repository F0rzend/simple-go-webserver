package domain

import (
	"net/mail"
	"time"
)

type User struct {
	ID        uint64
	Name      string
	Username  string
	Email     *mail.Address
	Balance   Balance
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(
	id uint64,
	name string,
	username string,
	email string,
	bitcoinAmount float64,
	usdBalance float64,
	createdAt time.Time,
	updatedAt time.Time,
) (*User, error) {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return nil, err
	}

	balance := NewBalance(
		BTCFromFloat(bitcoinAmount),
		USDFromFloat(usdBalance),
	)

	return &User{
		ID:        id,
		Name:      name,
		Username:  username,
		Email:     addr,
		Balance:   balance,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil

}
