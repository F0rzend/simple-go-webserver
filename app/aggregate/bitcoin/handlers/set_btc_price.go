package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

type SetBTCPriceRequest struct {
	Price float64 `json:"price"`
}

var ErrBadRequest = errors.New("validation error")

func (r SetBTCPriceRequest) Bind(_ *http.Request) error {
	if r.Price <= 0 {
		return ErrBadRequest
	}
	return nil
}

func (h *BitcoinHTTPHandlers) SetBTCPrice(w http.ResponseWriter, r *http.Request) {
	request := &SetBTCPriceRequest{}

	switch err := render.Bind(r, request); err {
	case nil:
	case ErrBadRequest:
		w.WriteHeader(http.StatusBadRequest)
		return
	default:
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.service.SetBTCPrice.Handle(service.SetBTCPrice{
		Price: request.Price,
	}); err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", "/bitcoin")
	render.Respond(w, r, nil)
}
