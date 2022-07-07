package service

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

type GetBTCPriceHandler struct {
	btcRepository entity.BTCRepository
}

func MustNewGetBTCCommand(btcRepository entity.BTCRepository) GetBTCPriceHandler {
	if btcRepository == nil {
		panic(ErrNilBTCRepository)
	}

	return GetBTCPriceHandler{
		btcRepository: btcRepository,
	}
}

func (h *GetBTCPriceHandler) Handle() entity.BTCPrice {
	return h.btcRepository.Get()
}
