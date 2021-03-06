package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	userService "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

type ChangeUSDBalanceRequest struct {
	Action string  `json:"action"`
	Amount float64 `json:"amount"`
}

var (
	ErrZeroAmount  = errors.New("amount can't be zero")
	ErrEmptyAction = errors.New("you should pass an action")
)

func (r ChangeUSDBalanceRequest) Bind(_ *http.Request) error {
	if r.Action == "" {
		return ErrEmptyAction
	}
	if r.Amount == 0 {
		return ErrZeroAmount
	}

	if _, err := bitcoinEntity.NewUSDAction(r.Action); err != nil {
		return err
	}

	return nil
}

func (h *UserHTTPHandlers) changeUSDBalance(w http.ResponseWriter, r *http.Request) {
	request := &ChangeUSDBalanceRequest{}

	id, err := getUserIDFromURL(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := render.Bind(r, request); err != nil {
		switch err {
		case ErrZeroAmount, ErrEmptyAction, bitcoinEntity.ErrInvalidAction:
			w.WriteHeader(http.StatusBadRequest)
		default:
			log.Error().Err(err).Send()
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = h.service.ChangeUSDBalance.Handle(userService.ChangeUSDBalance{
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)
}
