package handlers

import (
	"math/big"
	"net/http"
	"time"

	"github.com/F0rzend/simple-go-webserver/app/common"

	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

type UserResponse struct {
	ID         uint64     `json:"id"`
	Name       string     `json:"name"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	BTCBalance *big.Float `json:"btc_balance"`
	USDBalance *big.Float `json:"usd_balance"`
	CreatedAt  string     `json:"created_at"`
	UpdatedAt  string     `json:"updated_at"`
}

func UserToResponse(user *userEntity.User) *UserResponse {
	return &UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Username:   user.Username,
		Email:      user.Email.Address,
		BTCBalance: user.Balance.BTC.ToFloat(),
		USDBalance: user.Balance.USD.ToFloat(),
		CreatedAt:  user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  user.UpdatedAt.Format(time.RFC3339),
	}
}

func (h *UserHTTPHandlers) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := getUserIDFromURL(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.service.GetUser.Handle(id)
	switch err.(type) {
	case nil:
	case common.ErrUserNotFound:
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, UserToResponse(user))
}
