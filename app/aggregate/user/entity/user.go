package userentity

import (
	"fmt"
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
		return nil, common.NewFlaggedError("name cannot be empty", common.FlagInvalidArgument)
	}

	if username == "" {
		return nil, common.NewFlaggedError("username cannot be empty", common.FlagInvalidArgument)
	}

	addr, err := ParseEmail(email)
	if err != nil {
		return nil, fmt.Errorf("error parsing email: %w", err)
	}

	return &User{
		ID:        id,
		Name:      name,
		Username:  username,
		Email:     addr,
		Balance:   NewBalance(bitcoinentity.NewUSD(usdBalance), bitcoinentity.NewBTC(btcBalance)),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func ParseEmail(email string) (*mail.Address, error) {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return nil, common.NewFlaggedError("you must provide a valid email", common.FlagInvalidArgument)
	}
	return addr, nil
}

func (u *User) ChangeUSDBalance(action Action, amount bitcoinentity.USD) error {
	if amount.IsNegative() {
		return common.NewFlaggedError("amount cannot be negative", common.FlagInvalidArgument)
	}

	switch action { //nolint:exhaustive
	case DepositUSDAction:
		return u.deposit(amount)
	case WithdrawUSDAction:
		return u.withdraw(amount)
	default:
		return common.NewFlaggedError("invalid action", common.FlagInvalidArgument)
	}
}

func (u *User) deposit(amount bitcoinentity.USD) error {
	u.Balance.USD = u.Balance.USD.Add(amount)

	return nil
}

func (u *User) withdraw(amount bitcoinentity.USD) error {
	if u.Balance.USD.LessThan(amount) {
		return common.NewFlaggedError("the user does not have enough usd to withdraw", common.FlagInvalidArgument)
	}

	u.Balance.USD = u.Balance.USD.Sub(amount)

	return nil
}

func (u *User) ChangeBTCBalance(action Action, amount bitcoinentity.BTC, price bitcoinentity.BTCPrice) error {
	switch action { //nolint:exhaustive
	case BuyBTCAction:
		return u.buyBTC(amount, price)
	case SellBTCAction:
		return u.sellBTC(amount, price)
	default:
		return common.NewFlaggedError("invalid action", common.FlagInvalidArgument)
	}
}

func (u *User) buyBTC(amount bitcoinentity.BTC, price bitcoinentity.BTCPrice) error {
	if u.Balance.USD.LessThan(price.GetPrice()) {
		return common.NewFlaggedError("the user does not have enough usd to buy btc", common.FlagInvalidArgument)
	}

	u.Balance.USD = u.Balance.USD.Sub(amount.ToUSD(price))
	u.Balance.BTC = u.Balance.BTC.Add(amount)

	return nil
}

func (u *User) sellBTC(amount bitcoinentity.BTC, price bitcoinentity.BTCPrice) error {
	if u.Balance.BTC.LessThan(amount) {
		return common.NewFlaggedError("the user does not have enough btc to sell", common.FlagInvalidArgument)
	}

	u.Balance.BTC = u.Balance.BTC.Sub(amount)
	u.Balance.USD = u.Balance.USD.Add(amount.ToUSD(price))

	return nil
}
