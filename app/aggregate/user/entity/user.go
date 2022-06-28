package entity

import (
	"errors"
	"net/mail"
	"time"

	"github.com/F0rzend/simple-go-webserver/app/common"

	domain2 "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
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
	btcBalance float64,
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

	usdAmount, err := domain2.NewUSD(usdBalance)
	if err != nil {
		return nil, err
	}

	btcAmount, err := domain2.NewBTC(btcBalance)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        id,
		Name:      name,
		Username:  username,
		Email:     addr,
		Balance:   NewBalance(usdAmount, btcAmount),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func MustNewUser(
	id uint64,
	name string,
	username string,
	email string,
	btcBalance float64,
	usdBalance float64,
	createdAt time.Time,
	updatedAt time.Time,
) *User {
	user, err := NewUser(id, name, username, email, btcBalance, usdBalance, createdAt, updatedAt)
	if err != nil {
		panic(err)
	}

	return user
}

func (u *User) ChangeUSDBalance(action domain2.USDAction, amount domain2.USD) error {
	switch action {
	case domain2.DepositUSDAction:
		u.Balance.USD = u.Balance.USD.Add(amount)
	case domain2.WithdrawUSDAction:
		if u.Balance.USD.LessThan(amount) {
			return common.ErrInsufficientFunds(amount.ToFloat64())
		}
		updatedUSD, err := u.Balance.USD.Sub(amount)
		if err != nil {
			return err
		}
		u.Balance.USD = updatedUSD
	}

	return nil
}

func (u *User) ChangeBTCBalance(action domain2.BTCAction, amount domain2.BTC, price domain2.BTCPrice) error {
	switch action {
	case domain2.BuyBTCAction:
		if u.Balance.USD.LessThan(price.GetPrice()) {
			return common.ErrInsufficientFunds(amount.ToFloat64())
		}

		updatedUSD, err := u.Balance.USD.Sub(amount.ToUSD(price))
		if err != nil {
			return err
		}

		u.Balance.USD = updatedUSD
		u.Balance.BTC = u.Balance.BTC.Add(amount)
	case domain2.SellBTCAction:
		if u.Balance.BTC.LessThan(amount) {
			return common.ErrInsufficientFunds(amount.ToFloat64())
		}

		updatedBTC, err := u.Balance.BTC.Sub(amount)
		if err != nil {
			return err
		}

		u.Balance.BTC = updatedBTC
		u.Balance.USD = u.Balance.USD.Add(amount.ToUSD(price))
	default:
		return domain2.ErrInvalidAction
	}

	return nil
}
