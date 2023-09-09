package userhandlers

import (
	"net/http"

	"github.com/F0rzend/simple-go-webserver/pkg/hlog"

	"github.com/go-chi/render"

	"github.com/F0rzend/simple-go-webserver/app/common"
)

func (h *UserHTTPHandlers) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	logger := hlog.GetLoggerFromContext(r.Context())

	id, err := h.getUserIDFromRequest(r)
	if err != nil {
		logger.Error("failed to get user id from request", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	balance, err := h.service.GetUserBalance(id)
	if err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, map[string]any{"balance": balance.ToFloat()})
}
