package service

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
)

type SetBTCPrice struct {
	Price float64
}

type SetBTCPriceHandler struct {
	btcRepository entity.BTCRepository
}

func NewSetBTCPriceCommand(btcRepository entity.BTCRepository) (SetBTCPriceHandler, error) {
	if btcRepository == nil {
		return SetBTCPriceHandler{}, service.ErrNilBTCRepository
	}

	return SetBTCPriceHandler{
		btcRepository: btcRepository,
	}, nil
}

func MustNewSetBTCPriceCommand(btcRepository entity.BTCRepository) SetBTCPriceHandler {
	handler, err := NewSetBTCPriceCommand(btcRepository)
	if err != nil {
		panic(err)
	}

	return handler
}

func (h *SetBTCPriceHandler) Handle(command SetBTCPrice) error {
	price, err := entity.NewUSD(command.Price)
	if err != nil {
		return err
	}

	return h.btcRepository.SetPrice(price)
}
