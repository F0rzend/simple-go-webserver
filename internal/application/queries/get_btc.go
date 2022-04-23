package queries

import "github.com/F0rzend/SimpleGoWebserver/internal/domain"

type GetBTCQueryHandler struct {
	btcRepository domain.BTCRepository
}

func NewGetBTCCommand(btcRepository domain.BTCRepository) *GetBTCQueryHandler {
	return &GetBTCQueryHandler{btcRepository: btcRepository}
}

func (h *GetBTCQueryHandler) Handle() (domain.BTCPrice, error) {
	return h.btcRepository.Get()
}
