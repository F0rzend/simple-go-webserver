package service

import (
	"fmt"
	"net/http"

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
		return common.NewServiceError(http.StatusBadRequest, fmt.Sprintf("Invalid action: %s", cmd.Action))
	default:
		return err
	}

	switch err = us.userRepository.Update(cmd.UserID, func(user *userEntity.User) (*userEntity.User, error) {
		btc, err := bitcoinEntity.NewBTC(cmd.Amount)
		if err != nil {
			return nil, err
		}

		currentBitcoinPrice := us.bitcoinRepository.GetPrice()
		if err := user.ChangeBTCBalance(action, btc, currentBitcoinPrice); err != nil {
			return nil, err
		}
		return user, nil
	}); err {
	case repositories.ErrUserNotFound:
		return common.NewServiceError(
			http.StatusNotFound,
			fmt.Sprintf("User with id %d not found",
				cmd.UserID,
			),
		)
	case bitcoinEntity.ErrNegativeCurrency:
		return common.NewServiceError(
			http.StatusBadRequest,
			"The amount of currency cannot be negative. Please pass a number greater than 0",
		)
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
