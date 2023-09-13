package userhandlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	"github.com/F0rzend/simple-go-webserver/app/common"
)

type ChangeBTCBalanceRequest struct {
	Action string  `json:"action"`
	Amount float64 `json:"amount"`
}

func (r ChangeBTCBalanceRequest) Bind(_ *http.Request) error {
	if r.Action == "" {
		return common.NewValidationError("action cannot be empty")
	}
	if r.Amount == 0 {
		return common.NewValidationError("amount cannot be empty")
	}

	return nil
}

func (h *UserHTTPHandlers) ChangeBTCBalance(w http.ResponseWriter, r *http.Request) error {
	request := &ChangeBTCBalanceRequest{}

	id, err := h.getUserIDFromRequest(r)
	if err != nil {
		return fmt.Errorf("failed to get user id from request: %w", err)
	}

	if err := render.Bind(r, request); err != nil {
		return fmt.Errorf("failed to bind request: %w", err)
	}

	err = h.service.ChangeBitcoinBalance(id, request.Action, request.Amount)
	if err != nil {
		return common.NewValidationError(err.Error())
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)

	return nil
}
