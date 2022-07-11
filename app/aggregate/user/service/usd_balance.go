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

type ChangeUserBalanceCommand struct {
	UserID uint64
	Action string
	Amount float64
}

func (us *UserServiceImpl) ChangeUserBalance(cmd ChangeUserBalanceCommand) error {
	action, err := bitcoinEntity.NewUSDAction(cmd.Action)
	if err != nil {
		return common.NewServiceError(
			http.StatusBadRequest,
			fmt.Sprintf(
				"You must pass the correct action. Allowed: %s",
				strings.Join(bitcoinEntity.GetUSDActions(), ", "),
			),
		)
	}

	switch err := us.userRepository.Update(cmd.UserID, func(user *userEntity.User) (*userEntity.User, error) {
		usd, err := bitcoinEntity.NewUSD(cmd.Amount)
		if err != nil {
			return nil, err
		}

		if err := user.ChangeUSDBalance(action, usd); err != nil {
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
			"The user does not have enough funds to make a withdrawal",
		)
	default:
		return err
	}
}
