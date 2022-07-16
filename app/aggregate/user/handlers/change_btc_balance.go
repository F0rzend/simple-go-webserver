package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	userService "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

type ChangeBTCBalanceRequest struct {
	Action string  `json:"action"`
	Amount float64 `json:"amount"`
}

func (r ChangeBTCBalanceRequest) Bind(_ *http.Request) error {
	return nil
}

func (h *UserHTTPHandlers) ChangeBTCBalance(w http.ResponseWriter, r *http.Request) {
	request := &ChangeBTCBalanceRequest{}

	id, err := h.getUserIDFromRequest(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := render.Bind(r, request); err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	if err = h.service.ChangeBitcoinBalance(userService.ChangeBitcoinBalanceCommand{
		UserID: id,
		Action: request.Action,
		Amount: request.Amount,
	}); err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)
}
