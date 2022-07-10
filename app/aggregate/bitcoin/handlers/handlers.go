package handlers

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
)

type BitcoinHTTPHandlers struct {
	service *service.BitcoinService
}

func NewBitcoinHTTPHandlers(bitcoinService *service.BitcoinService) *BitcoinHTTPHandlers {
	return &BitcoinHTTPHandlers{
		service: bitcoinService,
	}
}
