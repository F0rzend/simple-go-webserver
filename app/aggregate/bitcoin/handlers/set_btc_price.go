package bitcoinhandlers

import (
	"fmt"
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/common"
	"github.com/go-chi/render"
)

type SetBTCPriceRequest struct {
	Price float64 `json:"price"`
}

func (r SetBTCPriceRequest) Bind(_ *http.Request) error {
	if r.Price == 0 {
		return common.NewValidationError("price is required")
	}

	return nil
}

func (h *BitcoinHTTPHandlers) SetBTCPrice(w http.ResponseWriter, r *http.Request) error {
	request := &SetBTCPriceRequest{}

	if err := render.Bind(r, request); err != nil {
		return err
	}

	err := h.service.SetBTCPrice(request.Price)
	if common.IsFlaggedError(err, common.FlagInvalidArgument) {
		return common.NewValidationError(err.Error())
	}
	if err != nil {
		return fmt.Errorf("failed to set btc price: %w", err)
	}

	render.Status(r, http.StatusNoContent)
	render.Respond(w, r, nil)

	return nil
}
