package queries

import "github.com/F0rzend/SimpleGoWebserver/internal/domain"

type GetUserBalanceQueryHandler struct {
	userRepository domain.UserRepository
	btcRepository  domain.BTCRepository
}

func NewGetUserBalanceQuery(
	userRepository domain.UserRepository,
	btcRepository domain.BTCRepository,
) *GetUserBalanceQueryHandler {
	return &GetUserBalanceQueryHandler{
		userRepository: userRepository,
		btcRepository:  btcRepository,
	}
}

func (h GetUserBalanceQueryHandler) Handle(userId uint64) (domain.USD, error) {
	user, err := h.userRepository.Get(userId)

	if err != nil {
		return 0, err
	}

	btc, err := h.btcRepository.Get()
	if err != nil {
		return 0, err
	}

	return user.Balance.Total(btc.Price), nil
}
