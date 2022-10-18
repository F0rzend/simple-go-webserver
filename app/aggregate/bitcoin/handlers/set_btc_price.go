package bitcoinhandlers

import (
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/common"
	"github.com/go-chi/render"
)

type SetBTCPriceRequest struct {
	Price float64 `json:"price"`
}

func (r SetBTCPriceRequest) Bind(_ *http.Request) error { return nil }

func (h *BitcoinHTTPHandlers) SetBTCPrice(w http.ResponseWriter, r *http.Request) {
	request := &SetBTCPriceRequest{}

	if err := render.Bind(r, request); err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	if err := h.service.SetBTCPrice(request.Price); err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	render.Status(r, http.StatusNoContent)
	render.Respond(w, r, nil)
}
