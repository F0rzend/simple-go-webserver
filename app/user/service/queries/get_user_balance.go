package queries

import (
	bitcoinDomain "github.com/F0rzend/simple-go-webserver/app/bitcoin/domain"
	userDomain "github.com/F0rzend/simple-go-webserver/app/user/domain"
)

type GetUserBalanceQueryHandler struct {
	userRepository userDomain.UserRepository
	btcRepository  bitcoinDomain.BTCRepository
}

func NewGetUserBalanceQuery(
	userRepository userDomain.UserRepository,
	btcRepository bitcoinDomain.BTCRepository,
) (GetUserBalanceQueryHandler, error) {
	if userRepository == nil {
		return GetUserBalanceQueryHandler{}, ErrNilUserRepository
	}
	if btcRepository == nil {
		return GetUserBalanceQueryHandler{}, ErrNilBTCRepository
	}

	return GetUserBalanceQueryHandler{
		userRepository: userRepository,
		btcRepository:  btcRepository,
	}, nil
}

func MustNewGetUserBalanceQuery(
	userRepository userDomain.UserRepository,
	btcRepository bitcoinDomain.BTCRepository,
) GetUserBalanceQueryHandler {
	handler, err := NewGetUserBalanceQuery(userRepository, btcRepository)
	if err != nil {
		panic(err)
	}

	return handler
}

func (h GetUserBalanceQueryHandler) Handle(userID uint64) (bitcoinDomain.USD, error) {
	user, err := h.userRepository.Get(userID)
	if err != nil {
		return bitcoinDomain.USD{}, err
	}

	return user.Balance.Total(h.btcRepository.Get()), nil
}
