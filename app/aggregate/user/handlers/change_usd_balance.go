package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	userService "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

type ChangeUSDBalanceRequest struct {
	Action string  `json:"action"`
	Amount float64 `json:"amount"`
}

var (
	ErrEmptyAction = common.NewApplicationError(
		http.StatusBadRequest,
		"Action cannot be empty",
	)
	ErrZeroAmount = common.NewApplicationError(
		http.StatusBadRequest,
		"Amount can't be zero",
	)
)

func (r ChangeUSDBalanceRequest) Bind(_ *http.Request) error {
	if r.Action == "" {
		return ErrEmptyAction
	}
	if r.Amount == 0 {
		return ErrZeroAmount
	}

	return nil
}

func (h *UserHTTPHandlers) ChangeUSDBalance(w http.ResponseWriter, r *http.Request) {
	request := &ChangeUSDBalanceRequest{}

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

	if err := h.service.ChangeUserBalance(userService.ChangeUserBalanceCommand{
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
