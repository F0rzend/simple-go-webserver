package domain

import (
	"errors"
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

var (
	ErrNegativeAmount    = errors.New("amount must be positive")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

func (u *User) ChangeUSDBalance(action USDAction, amount USD) error {
	if amount.IsNegative() {
		return ErrNegativeAmount
	}

	switch action {
	case DepositUSDAction:
		u.Balance.USD = u.Balance.USD.Add(amount)
	case WithdrawUSDAction:
		if u.Balance.USD < amount {
			return ErrInsufficientFunds
		}
		u.Balance.USD = u.Balance.USD.Sub(amount)
	default:
		return ErrInvalidAction
	}

	return nil
}

func (u *User) ChangeBTCBalance(action BTCAction, amount BTC, price USD) error {
	if amount.IsNegative() {
		return ErrNegativeAmount
	}

	switch action {
	case BuyBTCAction:
		print(u.Balance.USD.String(), " ", amount.ToUSD(price).String(), " ", price.String())
		if u.Balance.USD < amount.ToUSD(price) {
			return ErrInsufficientFunds
		}
		u.Balance.USD = u.Balance.USD.Sub(amount.ToUSD(price))
		u.Balance.BTC = u.Balance.BTC.Add(amount)
	case SellBTCAction:
		if u.Balance.BTC < amount {
			return ErrInsufficientFunds
		}
		u.Balance.BTC = u.Balance.BTC.Sub(amount)
		u.Balance.USD = u.Balance.USD.Add(amount.ToUSD(price))
	default:
		return ErrInvalidAction
	}

	return nil
}
