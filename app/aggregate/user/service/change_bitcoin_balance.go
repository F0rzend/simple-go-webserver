package service

import (
	"fmt"
	"net/http"
	"strings"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

type ChangeBitcoinBalanceCommand struct {
	UserID uint64
	Action string
	Amount float64
}

func (us *UserServiceImpl) ChangeBitcoinBalance(cmd ChangeBitcoinBalanceCommand) error {
	action, err := bitcoinEntity.NewBTCAction(cmd.Action)
	switch err {
	case nil:
	case bitcoinEntity.ErrInvalidAction:
		return common.NewServiceError(
			http.StatusBadRequest,
			fmt.Sprintf(
				"You must specify a valid action. Available actions: %s",
				strings.Join(bitcoinEntity.GetBTCActions(), ", "),
			),
		)
	default:
		return err
	}

	btc, err := bitcoinEntity.NewBTC(cmd.Amount)
	switch err {
	case nil:
	case bitcoinEntity.ErrNegativeCurrency:
		return common.NewServiceError(
			http.StatusBadRequest,
			"The amount of currency cannot be negative. Please pass a number greater than 0",
		)
	default:
		return err
	}

	user, err := us.userRepository.Get(cmd.UserID)
	switch err {
	case nil:
	case repositories.ErrUserNotFound:
		return common.NewServiceError(
			http.StatusNotFound,
			fmt.Sprintf("User with id %d not found",
				cmd.UserID,
			),
		)
	default:
		return err
	}

	currentBitcoinPrice := us.bitcoinRepository.GetPrice()

	switch err := user.ChangeBTCBalance(action, btc, currentBitcoinPrice); err {
	case nil:
		return us.userRepository.Save(user)
	case userEntity.ErrInsufficientFunds:
		return common.NewServiceError(
			http.StatusBadRequest,
			fmt.Sprintf(
				"The user does not have enough funds to %s BTC",
				cmd.Action,
			),
		)
	default:
		return err
	}
}
