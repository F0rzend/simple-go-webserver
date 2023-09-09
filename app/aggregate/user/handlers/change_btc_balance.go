package userhandlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	"github.com/F0rzend/simple-go-webserver/app/common"
	"github.com/F0rzend/simple-go-webserver/pkg/hlog"
)

type ChangeBTCBalanceRequest struct {
	Action string  `json:"action"`
	Amount float64 `json:"amount"`
}

func (r ChangeBTCBalanceRequest) Bind(_ *http.Request) error {
	if r.Action == "" {
		return ErrEmptyAction
	}
	if r.Amount == 0 {
		return ErrZeroAmount
	}

	return nil
}

func (h *UserHTTPHandlers) ChangeBTCBalance(w http.ResponseWriter, r *http.Request) {
	logger := hlog.GetLoggerFromContext(r.Context())

	request := &ChangeBTCBalanceRequest{}

	id, err := h.getUserIDFromRequest(r)
	if err != nil {
		logger.Error("failed to get user id from request", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := render.Bind(r, request); err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	if err = h.service.ChangeBitcoinBalance(id, request.Action, request.Amount); err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)
}
