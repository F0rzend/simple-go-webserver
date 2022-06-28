package handlers

import (
	"math/big"
	"net/http"

	userEntity "github.com/F0rzend/simple-go-webserver/app/common"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func (h *UserHTTPHandlers) getUserBalance(w http.ResponseWriter, r *http.Request) {
	id, err := getUserIDFromURL(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	balance, err := h.service.GetUserBalance.Handle(id)
	switch err.(type) {
	case nil:
	case userEntity.ErrUserNotFound:
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, map[string]*big.Float{"balance": balance.ToFloat()})
}
