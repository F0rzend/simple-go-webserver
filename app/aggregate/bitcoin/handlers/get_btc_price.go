package bitcoinhandlers

import (
	"math/big"
	"net/http"
	"time"

	"github.com/go-chi/render"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

type BTCResponse struct {
	Price     *big.Float `json:"price"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func BTCToResponse(btc bitcoinentity.BTCPrice) BTCResponse {
	return BTCResponse{
		Price:     btc.GetPrice().ToFloat(),
		UpdatedAt: btc.GetUpdatedAt(),
	}
}

func (h *BitcoinHTTPHandlers) GetBTCPrice(w http.ResponseWriter, r *http.Request) error {
	btc := h.service.GetBTCPrice()

	render.Status(r, http.StatusOK)
	render.Respond(w, r, map[string]BTCResponse{"btc": BTCToResponse(btc)})

	return nil
}
