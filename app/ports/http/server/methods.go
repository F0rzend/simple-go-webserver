package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	"github.com/F0rzend/simple-go-webserver/app/application/commands"
	"github.com/F0rzend/simple-go-webserver/app/domain"
	"github.com/F0rzend/simple-go-webserver/app/ports/http/types"
)

type hash map[string]any

func getUserIDFromURL(r *http.Request) (uint64, error) {
	return strconv.ParseUint(chi.URLParam(r, "id"), 10, 64) // nolint: gomnd
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request *types.CreateUserRequest

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

	id, err := s.app.Commands.CreateUser.Handle(commands.CreateUserCommand{
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

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := getUserIDFromURL(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := s.app.Queries.GetUser.Handle(id)
	switch err.(type) {
	case nil:
	case domain.ErrUserNotFound:
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, s.assembler.UserToResponse(user))
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var request *types.UpdateUserRequest
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

	err = s.app.Commands.UpdateUser.Handle(commands.UpdateUserCommand{
		ID:    id,
		Name:  request.Name,
		Email: request.Email,
	})
	switch err.(type) {
	case nil:
	case domain.ErrUserNotFound:
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

func (s *Server) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	id, err := getUserIDFromURL(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	balance, err := s.app.Queries.GetUserBalance.Handle(id)
	switch err.(type) {
	case nil:
	case domain.ErrUserNotFound:
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, hash{"balance": balance.ToFloat()})
}

func (s *Server) ChangeUSDBalance(w http.ResponseWriter, r *http.Request) {
	id, err := getUserIDFromURL(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var request *types.ChangeUSDBalanceRequest
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

	err = s.app.Commands.ChangeUSDBalance.Handle(commands.ChangeUSDBalanceCommand{
		UserID: id,
		Action: request.Action,
		Amount: request.Amount,
	})

	switch err.(type) {
	case nil:
	case domain.ErrInsufficientFunds, domain.ErrNegativeCurrency:
		w.WriteHeader(http.StatusBadRequest)
		return
	case domain.ErrUserNotFound:
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

func (s *Server) ChangeBTCBalance(w http.ResponseWriter, r *http.Request) {
	id, err := getUserIDFromURL(r)
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var request *types.ChangeBTCBalanceRequest
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

	err = s.app.Commands.ChangeBTCBalance.Handle(commands.ChangeBTCBalanceCommand{
		UserID: id,
		Action: request.Action,
		Amount: request.Amount,
	})
	switch err.(type) {
	case nil:
	case domain.ErrInsufficientFunds, domain.ErrNegativeCurrency:
		w.WriteHeader(http.StatusBadRequest)
		return
	case domain.ErrUserNotFound:
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

func (s *Server) GetBTC(w http.ResponseWriter, r *http.Request) {
	btc := s.app.Queries.GetBTC.Handle()

	render.Status(r, http.StatusOK)
	render.Respond(w, r, hash{"btc": s.assembler.BTCToResponse(btc)})
}

func (s *Server) SetBTCPrice(w http.ResponseWriter, r *http.Request) {
	var request *types.SetBTCPriceRequest
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

	if err := s.app.Commands.SetBTCPrice.Handle(commands.SetBTCPriceCommand{
		Price: request.Price,
	}); err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusNoContent)
	w.Header().Set("Location", "/bitcoin")
	render.Respond(w, r, nil)
}
