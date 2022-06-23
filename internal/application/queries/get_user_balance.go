package queries

import "github.com/F0rzend/simple-go-webserver/internal/domain"

type GetUserBalanceQueryHandler struct {
	userRepository domain.UserRepository
	btcRepository  domain.BTCRepository
}

func NewGetUserBalanceQuery(
	userRepository domain.UserRepository,
	btcRepository domain.BTCRepository,
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
	userRepository domain.UserRepository,
	btcRepository domain.BTCRepository,
) GetUserBalanceQueryHandler {
	handler, err := NewGetUserBalanceQuery(userRepository, btcRepository)
	if err != nil {
		panic(err)
	}

	return handler
}

func (h GetUserBalanceQueryHandler) Handle(userID uint64) (domain.USD, error) {
	user, err := h.userRepository.Get(userID)
	if err != nil {
		return domain.USD{}, err
	}

	return user.Balance.Total(h.btcRepository.Get()), nil
}
