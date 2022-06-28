package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	userService "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

type ChangeBTCBalanceRequest struct {
	Action string  `json:"action"`
	Amount float64 `json:"amount"`
}

func (r ChangeBTCBalanceRequest) Bind(_ *http.Request) error {
	if _, err := bitcoinEntity.NewBTCAction(r.Action); err != nil {
		return err
	}

	return nil
}

func (h *UserHTTPHandlers) changeBTCBalance(w http.ResponseWriter, r *http.Request) {
	request := &ChangeBTCBalanceRequest{}

	id, err := getUserIDFromURL(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch err := render.Bind(r, request); err {
	case nil:
	case bitcoinEntity.ErrInvalidAction:
		w.WriteHeader(http.StatusBadRequest)
		return
	default:
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.service.ChangeBTCBalance.Handle(userService.ChangeBTCBalance{
		UserID: id,
		Action: request.Action,
		Amount: request.Amount,
	})
	switch err.(type) {
	case nil:
	case common.ErrInsufficientFunds, common.ErrNegativeCurrency:
		w.WriteHeader(http.StatusBadRequest)
		return
	case common.ErrUserNotFound:
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)
}
