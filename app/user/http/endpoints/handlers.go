package endpoints

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	bitcoinDomain "github.com/F0rzend/simple-go-webserver/app/bitcoin/domain"
	userDomain "github.com/F0rzend/simple-go-webserver/app/user/domain"
	"github.com/F0rzend/simple-go-webserver/app/user/http/types"
	"github.com/F0rzend/simple-go-webserver/app/user/service/commands"
)

func getUserIDFromURL(r *http.Request) (uint64, error) {
	return strconv.ParseUint(chi.URLParam(r, "id"), 10, 64) // nolint: gomnd
}

func (u *UserHTTPEndpoints) CreateUser(w http.ResponseWriter, r *http.Request) {
	request := &types.CreateUserRequest{}

	if err := render.Bind(r, request); err != nil {
		switch err {
		case types.ErrBadRequest:
			w.WriteHeader(http.StatusBadRequest)
		default:
			log.Error().Err(err).Send()
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	id, err := u.service.Commands.CreateUser.Handle(commands.CreateUserCommand{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
	})
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)
}

func (u *UserHTTPEndpoints) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := getUserIDFromURL(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := u.service.Queries.GetUser.Handle(id)
	switch err.(type) {
	case nil:
	case userDomain.ErrUserNotFound:
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, types.UserToResponse(user))
}

func (u *UserHTTPEndpoints) UpdateUser(w http.ResponseWriter, r *http.Request) {
	request := &types.UpdateUserRequest{}

	id, err := getUserIDFromURL(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := render.Bind(r, request); err != nil {
		switch err {
		case types.ErrBadRequest:
			w.WriteHeader(http.StatusBadRequest)
		default:
			log.Error().Err(err).Send()
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = u.service.Commands.UpdateUser.Handle(commands.UpdateUserCommand{
		ID:    id,
		Name:  request.Name,
		Email: request.Email,
	})
	switch err.(type) {
	case nil:
	case userDomain.ErrUserNotFound:
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)
}

func (u *UserHTTPEndpoints) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	id, err := getUserIDFromURL(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	balance, err := u.service.Queries.GetUserBalance.Handle(id)
	switch err.(type) {
	case nil:
	case userDomain.ErrUserNotFound:
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, map[string]*big.Float{"balance": balance.ToFloat()})
}

func (u *UserHTTPEndpoints) ChangeUSDBalance(w http.ResponseWriter, r *http.Request) {
	id, err := getUserIDFromURL(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := &types.ChangeUSDBalanceRequest{}

	if err := render.Bind(r, request); err != nil {
		switch err {
		case types.ErrBadRequest:
			w.WriteHeader(http.StatusBadRequest)
		default:
			log.Error().Err(err).Send()
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = u.service.Commands.ChangeUSDBalance.Handle(commands.ChangeUSDBalanceCommand{
		UserID: id,
		Action: request.Action,
		Amount: request.Amount,
	})

	switch err.(type) {
	case nil:
	case userDomain.ErrInsufficientFunds, bitcoinDomain.ErrNegativeCurrency:
		w.WriteHeader(http.StatusBadRequest)
		return
	case userDomain.ErrUserNotFound:
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)
}

func (u *UserHTTPEndpoints) ChangeBTCBalance(w http.ResponseWriter, r *http.Request) {
	id, err := getUserIDFromURL(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := &types.ChangeBTCBalanceRequest{}

	if err := render.Bind(r, request); err != nil {
		switch err {
		case types.ErrBadRequest:
			w.WriteHeader(http.StatusBadRequest)
		default:
			log.Error().Err(err).Send()
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = u.service.Commands.ChangeBTCBalance.Handle(commands.ChangeBTCBalanceCommand{
		UserID: id,
		Action: request.Action,
		Amount: request.Amount,
	})
	switch err.(type) {
	case nil:
	case userDomain.ErrInsufficientFunds, bitcoinDomain.ErrNegativeCurrency:
		w.WriteHeader(http.StatusBadRequest)
		return
	case userDomain.ErrUserNotFound:
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	render.Respond(w, r, nil)
}
