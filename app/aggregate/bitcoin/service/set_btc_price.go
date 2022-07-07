package service

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

type SetBTCPrice struct {
	Price float64
}

type SetBTCPriceHandler struct {
	btcRepository entity.BTCRepository
}

func MustNewSetBTCPriceCommand(btcRepository entity.BTCRepository) SetBTCPriceHandler {
	if btcRepository == nil {
		panic(ErrNilBTCRepository)
	}

	return SetBTCPriceHandler{
		btcRepository: btcRepository,
	}
}

func (h *SetBTCPriceHandler) Handle(command SetBTCPrice) error {
	price, err := entity.NewUSD(command.Price)
	if err != nil {
		return err
	}

	return h.btcRepository.SetPrice(price)
}
