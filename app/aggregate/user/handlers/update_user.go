package userhandlers

import (
	"fmt"
	"net/http"

	"github.com/F0rzend/simple-go-webserver/pkg/hlog"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/common"
	"github.com/go-chi/render"
)

type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

var ErrNothingToUpdate = common.NewApplicationError(
	http.StatusBadRequest,
	"At least one field must be updated",
)

func (r UpdateUserRequest) Bind(_ *http.Request) error {
	if r.Name == nil && r.Email == nil {
		return ErrNothingToUpdate
	}

	if r.Email != nil {
		if _, err := userentity.ParseEmail(*r.Email); err != nil {
			return err
		}
	}
	return nil
}

func (h *UserHTTPHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	logger := hlog.GetLoggerFromContext(r.Context())

	request := &UpdateUserRequest{}

	id, err := h.getUserIDFromRequest(r)
	if err != nil {
		logger.Error("Error while getting user id from request", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := render.Bind(r, request); err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	if err := h.service.UpdateUser(id, request.Name, request.Email); err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)
}
