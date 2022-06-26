package endpoints

import (
	"github.com/F0rzend/simple-go-webserver/app/bitcoin/service"
	"github.com/go-chi/chi/v5"
)

type BitcoinHTTPEndpoints struct {
	service *service.BitcoinService
}

func NewBitcoinHTTPEndpoints(service *service.BitcoinService) *BitcoinHTTPEndpoints {
	return &BitcoinHTTPEndpoints{
		service: service,
	}
}

func (b *BitcoinHTTPEndpoints) Register(r chi.Router) {
	r.Route("/bitcoin", func(r chi.Router) {
		r.Get("/", b.GetBTC)
		r.Put("/", b.SetBTCPrice)
	})
}
