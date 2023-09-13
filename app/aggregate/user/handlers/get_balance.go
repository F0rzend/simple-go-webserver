package userhandlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	"github.com/F0rzend/simple-go-webserver/app/common"
)

func (h *UserHTTPHandlers) GetUserBalance(w http.ResponseWriter, r *http.Request) error {
	id, err := h.getUserIDFromRequest(r)
	if err != nil {
		return fmt.Errorf("failed to get user id from request: %w", err)
	}

	balance, err := h.service.GetUserBalance(id)
	if common.IsFlaggedError(err, common.FlagNotFound) {
		return common.NewNotFoundError("user not found")
	}
	if err != nil {
		return fmt.Errorf("failed to get user balance: %w", err)
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, map[string]any{"balance": balance.ToFloat()})

	return nil
}
