package queries

import (
	"github.com/F0rzend/simple-go-webserver/app/domain"
)

type GetBTCQueryHandler struct {
	btcRepository domain.BTCRepository
}

func NewGetBTCCommand(btcRepository domain.BTCRepository) (GetBTCQueryHandler, error) {
	if btcRepository == nil {
		return GetBTCQueryHandler{}, ErrNilBTCRepository
	}

	return GetBTCQueryHandler{
		btcRepository: btcRepository,
	}, nil
}

func MustNewGetBTCCommand(btcRepository domain.BTCRepository) GetBTCQueryHandler {
	cmd, err := NewGetBTCCommand(btcRepository)
	if err != nil {
		panic(err)
	}

	return cmd
}

func (h *GetBTCQueryHandler) Handle() domain.BTCPrice {
	return h.btcRepository.Get()
}
