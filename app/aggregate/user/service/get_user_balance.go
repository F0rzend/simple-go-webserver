package service

import (
	"fmt"
	"net/http"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

type GetUserBalanceHandler struct {
	userRepository userEntity.UserRepository
	btcRepository  bitcoinEntity.BTCRepository
}

func MustNewGetUserBalance(
	userRepository userEntity.UserRepository,
	btcRepository bitcoinEntity.BTCRepository,
) GetUserBalanceHandler {
	if userRepository == nil {
		panic(ErrNilUserRepository)
	}
	if btcRepository == nil {
		panic(ErrNilBTCRepository)
	}

	return GetUserBalanceHandler{
		userRepository: userRepository,
		btcRepository:  btcRepository,
	}
}

func (h GetUserBalanceHandler) Handle(userID uint64) (bitcoinEntity.USD, error) {
	user, err := h.userRepository.Get(userID)
	switch err {
	case nil:
		return user.Balance.Total(h.btcRepository.Get()), nil
	case userRepositories.ErrUserNotFound:
		return bitcoinEntity.USD{}, common.NewServiceError(
			http.StatusNotFound,
			fmt.Sprintf(
				"User with id %d not found",
				userID,
			),
		)
	default:
		return bitcoinEntity.USD{}, err
	}
}
