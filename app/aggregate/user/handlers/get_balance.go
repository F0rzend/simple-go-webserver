package handlers

import (
	"math/big"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	"github.com/F0rzend/simple-go-webserver/app/common"
)

func (h *UserHTTPHandlers) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	id, err := getUserIDFromURL(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	balance, err := h.service.GetUserBalance(id)
	if err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, map[string]*big.Float{"balance": balance.ToFloat()})
}
