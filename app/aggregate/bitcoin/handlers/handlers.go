package bitcoinhandlers

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

type BitcoinHTTPHandlers struct {
	service BitcoinService
}

func NewBitcoinHTTPHandlers(bitcoinService BitcoinService) *BitcoinHTTPHandlers {
	return &BitcoinHTTPHandlers{
		service: bitcoinService,
	}
}

//go:generate moq -out "mock_bitcoin_service.gen.go" . BitcoinService:MockBitcoinService
type BitcoinService interface {
	GetBTCPrice() bitcoinentity.BTCPrice
	SetBTCPrice(newPrice float64) error
}
