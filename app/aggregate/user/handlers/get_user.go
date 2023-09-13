package userhandlers

import (
	"fmt"
	"math/big"
	"net/http"
	"time"

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

func (h *UserHTTPHandlers) GetUser(w http.ResponseWriter, r *http.Request) error {
	id, err := h.getUserIDFromRequest(r)
	if err != nil {
		return fmt.Errorf("failed to get user id from request: %w", err)
	}

	user, err := h.service.GetUser(id)
	if common.IsFlaggedError(err, common.FlagNotFound) {
		return common.NewNotFoundError("user not found")
	}
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, UserToResponse(user))

	return nil
}
