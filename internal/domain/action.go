package domain

import "errors"

var (
	DepositUSDAction  = USDAction{"deposit"}
	WithdrawUSDAction = USDAction{"withdraw"}

	BuyBTCAction  = BTCAction{"buy"}
	SellBTCAction = BTCAction{"sell"}
)

var (
	usdActions = []USDAction{
		DepositUSDAction,
		WithdrawUSDAction,
	}
	btcActions = []BTCAction{
		BuyBTCAction,
		SellBTCAction,
	}
)

type (
	USDAction action
	BTCAction action
)

type action struct {
	action string
}

var ErrInvalidAction = errors.New("invalid action")

func NewUSDAction(action string) (USDAction, error) {
	for _, a := range usdActions {
		if a.action == action {
			return a, nil
		}
	}
	return USDAction{}, ErrInvalidAction
}

func NewBTCAction(action string) (BTCAction, error) {
	for _, a := range btcActions {
		if a.action == action {
			return a, nil
		}
	}
	return BTCAction{}, ErrInvalidAction
}
