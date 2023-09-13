package userhandlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	"github.com/F0rzend/simple-go-webserver/app/common"
)

type ChangeUSDBalanceRequest struct {
	Action string  `json:"action"`
	Amount float64 `json:"amount"`
}

func (r ChangeUSDBalanceRequest) Bind(_ *http.Request) error {
	if r.Action == "" {
		return common.NewValidationError("action is required")
	}
	if r.Amount == 0 {
		return common.NewValidationError("amount is required")
	}

	return nil
}

func (h *UserHTTPHandlers) ChangeUSDBalance(w http.ResponseWriter, r *http.Request) error {
	request := &ChangeUSDBalanceRequest{}

	id, err := h.getUserIDFromRequest(r)
	if err != nil {
		return fmt.Errorf("failed to get user id from request: %w", err)
	}

	if err := render.Bind(r, request); err != nil {
		return fmt.Errorf("failed to bind request: %w", err)
	}

	err = h.service.ChangeUserBalance(id, request.Action, request.Amount)
	if common.IsFlaggedError(err, common.FlagInvalidArgument) {
		return common.NewValidationError(err.Error())
	}
	if err != nil {
		return fmt.Errorf("failed to change user balance: %w", err)
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)

	return nil
}
