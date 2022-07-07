package service

import (
	"fmt"
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/common"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type ChangeBTCBalance struct {
	UserID uint64
	Action string
	Amount float64
}

type ChangeBTCBalanceHandler struct {
	userRepository userEntity.UserRepository
	btcRepository  bitcoinEntity.BTCRepository
}

func MustNewChangeBTCBalanceHandler(
	userRepository userEntity.UserRepository,
	btcRepository bitcoinEntity.BTCRepository,
) ChangeBTCBalanceHandler {
	if userRepository == nil {
		panic(ErrNilUserRepository)
	}
	if btcRepository == nil {
		panic(ErrNilBTCRepository)
	}

	return ChangeBTCBalanceHandler{
		userRepository: userRepository,
		btcRepository:  btcRepository,
	}
}

func (h ChangeBTCBalanceHandler) Handle(cmd ChangeBTCBalance) error {
	action, err := bitcoinEntity.NewBTCAction(cmd.Action)
	if err != nil {
		return err
	}

	switch err = h.userRepository.Update(cmd.UserID, func(user *userEntity.User) (*userEntity.User, error) {
		btc, err := bitcoinEntity.NewBTC(cmd.Amount)
		if err != nil {
			return nil, err
		}

		if err := user.ChangeBTCBalance(action, btc, h.btcRepository.Get()); err != nil {
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
