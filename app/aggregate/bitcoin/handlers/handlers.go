package handlers

import (
	"github.com/go-chi/chi/v5"

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

func (h *BitcoinHTTPHandlers) SetRoutes(r chi.Router) {
	const bitcoinPath = "/bitcoin"

	r.Get(bitcoinPath, h.getBTCPrice)
	r.Put(bitcoinPath, h.setBTCPrice)
}
