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

var (
	ErrNameEmpty     = errors.New("name is empty")
	ErrUsernameEmpty = errors.New("username is empty")
)

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
	if name == "" {
		return nil, ErrNameEmpty
	}

	if username == "" {
		return nil, ErrUsernameEmpty
	}

	addr, err := mail.ParseAddress(email)
	if err != nil {
		return nil, err
	}

	btc, err := BTCFromFloat(bitcoinAmount)
	if err != nil {
		return nil, err
	}
	usd, err := USDFromFloat(usdBalance)
	if err != nil {
		return nil, err
	}
	balance := NewBalance(usd, btc)

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

var ErrInsufficientFunds = errors.New("insufficient funds")

func (u *User) ChangeUSDBalance(action USDAction, amount USD) error {
	switch action {
	case DepositUSDAction:
		updatedUSD, err := u.Balance.USD.Add(amount)
		if err != nil {
			return err
		}
		u.Balance.USD = updatedUSD
	case WithdrawUSDAction:
		if u.Balance.USD.GetCent() < amount.GetCent() {
			return ErrInsufficientFunds
		}
		updatedUSD, err := u.Balance.USD.Sub(amount)
		if err != nil {
			return err
		}
		u.Balance.USD = updatedUSD
	}

	return nil
}

func (u *User) ChangeBTCBalance(action BTCAction, amount BTC, price BTCPrice) error {
	switch action {
	case BuyBTCAction:
		if u.Balance.USD.GetCent() < amount.ToUSD(price).GetCent() {
			return ErrInsufficientFunds
		}

		updatedUSD, err := u.Balance.USD.Sub(amount.ToUSD(price))
		if err != nil {
			return err
		}

		updatedBTC, err := u.Balance.BTC.Add(amount)
		if err != nil {
			return err
		}

		u.Balance.USD = updatedUSD
		u.Balance.BTC = updatedBTC
	case SellBTCAction:
		if u.Balance.BTC.GetSatoshi() < amount.GetSatoshi() {
			return ErrInsufficientFunds
		}

		updatedBTC, err := u.Balance.BTC.Sub(amount)
		if err != nil {
			return err
		}

		updatedUSD, err := u.Balance.USD.Add(amount.ToUSD(price))
		if err != nil {
			return err
		}

		u.Balance.BTC = updatedBTC
		u.Balance.USD = updatedUSD
	default:
		return ErrInvalidAction
	}

	return nil
}
