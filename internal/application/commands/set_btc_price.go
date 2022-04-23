package commands

import "github.com/F0rzend/SimpleGoWebserver/internal/domain"

type SetBTCPriceCommand struct {
	Price float64
}

type SetBTCPriceCommandHandler struct {
	btcRepository domain.BTCRepository
}

func NewSetBTCPriceCommand(btcRepository domain.BTCRepository) *SetBTCPriceCommandHandler {
	return &SetBTCPriceCommandHandler{btcRepository: btcRepository}
}

func (h *SetBTCPriceCommandHandler) Handle(command SetBTCPriceCommand) error {
	return h.btcRepository.SetPrice(domain.NewUSD(command.Price))
}
