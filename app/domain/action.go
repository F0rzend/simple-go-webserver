package domain

import "errors"

var (
	DepositUSDAction  = USDAction{"deposit"}
	WithdrawUSDAction = USDAction{"withdraw"}

	BuyBTCAction  = BTCAction{"buy"}
	SellBTCAction = BTCAction{"sell"}
)

var (
	usdActions = map[string]USDAction{
		"deposit":  DepositUSDAction,
		"withdraw": WithdrawUSDAction,
	}
	btcActions = map[string]BTCAction{
		"buy":  BuyBTCAction,
		"sell": SellBTCAction,
	}
)

type (
	USDAction action
	BTCAction action
)

type action struct {
	a string
}

func (a action) String() string {
	return a.a
}

var ErrInvalidAction = errors.New("invalid action")

func NewUSDAction(action string) (USDAction, error) {
	usdAction, ok := usdActions[action]
	if !ok {
		return USDAction{}, ErrInvalidAction
	}

	return usdAction, nil
}

func NewBTCAction(action string) (BTCAction, error) {
	btcAction, ok := btcActions[action]
	if !ok {
		return BTCAction{}, ErrInvalidAction
	}
	return btcAction, nil
}
