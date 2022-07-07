package entity

import (
	"errors"
	"net/mail"
	"time"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
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
	ErrInvalidEmail  = errors.New("invalid email")
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
		return nil, ErrInvalidEmail
	}

	usdAmount, err := bitcoinEntity.NewUSD(usdBalance)
	if err != nil {
		return nil, err
	}

	btcAmount, err := bitcoinEntity.NewBTC(btcBalance)
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

func (u *User) ChangeUSDBalance(action bitcoinEntity.USDAction, amount bitcoinEntity.USD) error {
	switch action {
	case bitcoinEntity.DepositUSDAction:
		u.Balance.USD = u.Balance.USD.Add(amount)
	case bitcoinEntity.WithdrawUSDAction:
		if u.Balance.USD.LessThan(amount) {
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

func (u *User) ChangeBTCBalance(action bitcoinEntity.BTCAction, amount bitcoinEntity.BTC, price bitcoinEntity.BTCPrice) error {
	switch action {
	case bitcoinEntity.BuyBTCAction:
		if u.Balance.USD.LessThan(price.GetPrice()) {
			return ErrInsufficientFunds
		}

		updatedUSD, err := u.Balance.USD.Sub(amount.ToUSD(price))
		if err != nil {
			return err
		}

		u.Balance.USD = updatedUSD
		u.Balance.BTC = u.Balance.BTC.Add(amount)
	case bitcoinEntity.SellBTCAction:
		if u.Balance.BTC.LessThan(amount) {
			return ErrInsufficientFunds
		}

		updatedBTC, err := u.Balance.BTC.Sub(amount)
		if err != nil {
			return err
		}

		u.Balance.BTC = updatedBTC
		u.Balance.USD = u.Balance.USD.Add(amount.ToUSD(price))
	default:
		return bitcoinEntity.ErrInvalidAction
	}

	return nil
}
