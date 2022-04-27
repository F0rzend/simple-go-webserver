package commands

import (
	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
)

type SetBTCPriceCommand struct {
	Price float64
}

type SetBTCPriceCommandHandler struct {
	btcRepository domain.BTCRepository
}

func NewSetBTCPriceCommand(btcRepository domain.BTCRepository) (SetBTCPriceCommandHandler, error) {
	if btcRepository == nil {
		return SetBTCPriceCommandHandler{}, ErrNilBTCRepository
	}

	return SetBTCPriceCommandHandler{
		btcRepository: btcRepository,
	}, nil
}

func MustNewSetBTCPriceCommand(btcRepository domain.BTCRepository) SetBTCPriceCommandHandler {
	handler, err := NewSetBTCPriceCommand(btcRepository)
	if err != nil {
		panic(err)
	}

	return handler
}

func (h *SetBTCPriceCommandHandler) Handle(command SetBTCPriceCommand) error {
	price, err := domain.USDFromFloat(command.Price)
	if err != nil {
		return err
	}

	return h.btcRepository.SetPrice(price)
}
