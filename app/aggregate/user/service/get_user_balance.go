package service

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	userDomain "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type GetUserBalanceHandler struct {
	userRepository userDomain.UserRepository
	btcRepository  entity.BTCRepository
}

func NewGetUserBalance(
	userRepository userDomain.UserRepository,
	btcRepository entity.BTCRepository,
) (GetUserBalanceHandler, error) {
	if userRepository == nil {
		return GetUserBalanceHandler{}, ErrNilUserRepository
	}
	if btcRepository == nil {
		return GetUserBalanceHandler{}, ErrNilBTCRepository
	}

	return GetUserBalanceHandler{
		userRepository: userRepository,
		btcRepository:  btcRepository,
	}, nil
}

func MustNewGetUserBalance(
	userRepository userDomain.UserRepository,
	btcRepository entity.BTCRepository,
) GetUserBalanceHandler {
	handler, err := NewGetUserBalance(userRepository, btcRepository)
	if err != nil {
		panic(err)
	}

	return handler
}

func (h GetUserBalanceHandler) Handle(userID uint64) (entity.USD, error) {
	user, err := h.userRepository.Get(userID)
	if err != nil {
		return entity.USD{}, err
	}

	return user.Balance.Total(h.btcRepository.Get()), nil
}
