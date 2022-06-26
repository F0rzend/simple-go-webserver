package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	"github.com/F0rzend/simple-go-webserver/app/bitcoin/http/types"
	"github.com/F0rzend/simple-go-webserver/app/bitcoin/service/commands"
)

func (b *BitcoinHTTPEndpoints) GetBTC(w http.ResponseWriter, r *http.Request) {
	btc := b.service.Queries.GetBTC.Handle()

	render.Status(r, http.StatusOK)
	render.Respond(w, r, map[string]types.BTCResponse{"btc": types.BTCToResponse(btc)})
}

func (b *BitcoinHTTPEndpoints) SetBTCPrice(w http.ResponseWriter, r *http.Request) {
	request := &types.SetBTCPriceRequest{}

	if err := render.Bind(r, request); err != nil {
		switch err {
		case types.ErrBadRequest:
			w.WriteHeader(http.StatusBadRequest)
		default:
			log.Error().Err(err).Send()
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	if err := b.service.Commands.SetBTCPrice.Handle(commands.SetBTCPriceCommand{
		Price: request.Price,
	}); err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", "/bitcoin")
	render.Respond(w, r, nil)
}
