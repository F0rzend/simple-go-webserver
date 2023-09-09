package userhandlers

import (
	"math/big"
	"net/http"
	"time"

	"github.com/F0rzend/simple-go-webserver/pkg/hlog"

	"github.com/go-chi/render"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

type UserResponse struct {
	ID         uint64     `json:"id"`
	Name       string     `json:"name"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	BTCBalance *big.Float `json:"btc_balance"`
	USDBalance *big.Float `json:"usd_balance"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func UserToResponse(user *userentity.User) *UserResponse {
	return &UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Username:   user.Username,
		Email:      user.Email.Address,
		BTCBalance: user.Balance.BTC.ToFloat(),
		USDBalance: user.Balance.USD.ToFloat(),
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}

func (h *UserHTTPHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	logger := hlog.GetLoggerFromContext(r.Context())

	id, err := h.getUserIDFromRequest(r)
	if err != nil {
		logger.Error("error getting user id from request", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		common.RenderHTTPError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, UserToResponse(user))
}
