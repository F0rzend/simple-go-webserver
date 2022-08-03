package userentity

import (
	"net/http"
	"net/mail"
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/common"
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
	ErrNameEmpty = common.NewApplicationError(
		http.StatusBadRequest,
		"Name cannot be empty",
	)
	ErrUsernameEmpty = common.NewApplicationError(
		http.StatusBadRequest,
		"Username cannot be empty",
	)
	ErrInvalidEmail = common.NewApplicationError(
		http.StatusBadRequest,
		"You must provide a valid email",
	)
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

	addr, err := ParseEmail(email)
	if err != nil {
		return nil, err
	}

	usdAmount, err := bitcoinentity.NewUSD(usdBalance)
	if err != nil {
		return nil, err
	}

	btcAmount, err := bitcoinentity.NewBTC(btcBalance)
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

func ParseEmail(email string) (*mail.Address, error) {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return nil, ErrInvalidEmail
	}
	return addr, nil
}

func (u *User) ChangeUSDBalance(action Action, amount bitcoinentity.USD) error {
	switch action {
	case DepositUSDAction:
		return u.deposit(amount)
	case WithdrawUSDAction:
		return u.withdraw(amount)
	default:
		return ErrInvalidUSDAction
	}
}

func (u *User) deposit(amount bitcoinentity.USD) error {
	u.Balance.USD = u.Balance.USD.Add(amount)

	return nil
}

func (u *User) withdraw(amount bitcoinentity.USD) error {
	if u.Balance.USD.LessThan(amount) {
		return ErrInsufficientFunds
	}

	updatedUSD, err := u.Balance.USD.Sub(amount)
	if err != nil {
		return err
	}
	u.Balance.USD = updatedUSD

	return nil
}

func (u *User) ChangeBTCBalance(action Action, amount bitcoinentity.BTC, price bitcoinentity.BTCPrice) error {
	switch action {
	case BuyBTCAction:
		return u.buyBTC(amount, price)
	case SellBTCAction:
		return u.sellBTC(amount, price)
	default:
		return ErrInvalidBTCAction
	}
}

func (u *User) buyBTC(amount bitcoinentity.BTC, price bitcoinentity.BTCPrice) error {
	if u.Balance.USD.LessThan(price.GetPrice()) {
		return ErrInsufficientFunds
	}

	updatedUSD, err := u.Balance.USD.Sub(amount.ToUSD(price))
	if err != nil {
		return err
	}

	u.Balance.USD = updatedUSD
	u.Balance.BTC = u.Balance.BTC.Add(amount)

	return nil
}

func (u *User) sellBTC(amount bitcoinentity.BTC, price bitcoinentity.BTCPrice) error {
	if u.Balance.BTC.LessThan(amount) {
		return ErrInsufficientFunds
	}

	updatedBTC, err := u.Balance.BTC.Sub(amount)
	if err != nil {
		return err
	}

	u.Balance.BTC = updatedBTC
	u.Balance.USD = u.Balance.USD.Add(amount.ToUSD(price))

	return nil
}
