package userentity

import (
	"fmt"
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/common"
)

var (
	ErrInsufficientFunds = common.NewApplicationError(
		http.StatusBadRequest,
		"The user does not have enough funds",
	)
	ErrInvalidUSDAction = common.NewApplicationError(
		http.StatusBadRequest,
		fmt.Sprintf(
			"You must specify a valid action. Available actions: %s and %s",
			DepositUSDAction, WithdrawUSDAction,
		),
	)
	ErrInvalidBTCAction = common.NewApplicationError(
		http.StatusBadRequest,
		fmt.Sprintf(
			"You must specify a valid action. Available actions: %s and %s",
			BuyBTCAction, SellBTCAction,
		),
	)
)
