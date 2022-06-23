package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/F0rzend/simple-go-webserver/internal/domain"
	"github.com/go-chi/chi/v5"

	"github.com/go-chi/render"

	"github.com/F0rzend/simple-go-webserver/internal/application/commands"
	"github.com/F0rzend/simple-go-webserver/internal/ports/http/types"
)

func getUserIDFromURL(r *http.Request) (uint64, error) {
	return strconv.ParseUint(chi.URLParam(r, "id"), 10, 64) // nolint: gomnd
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	request := new(types.CreateUserRequest)

	if err := render.Bind(r, request); err != nil {
		switch err {
		case types.ErrBadRequest:
			rspd.StatusOnly(http.StatusBadRequest)
		default:
			rspd.InternalError(err)
		}
		return
	}

	id, err := s.app.Commands.CreateUser.Handle(commands.CreateUserCommand{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
	})
	if err != nil {
		rspd.InternalError(err)
		return
	}

	rspd.Status(http.StatusNoContent)
	rspd.LocationHeader(fmt.Sprintf("/users/%d", id))
	rspd.Response(nil)
}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	id, err := getUserIDFromURL(r)
	if err != nil {
		rspd.InternalError(err)
		return
	}

	user, err := s.app.Queries.GetUser.Handle(id)
	switch err.(type) {
	case nil:
	case domain.ErrUserNotFound:
		rspd.StatusOnly(http.StatusNotFound)
		return
	default:
		rspd.InternalError(err)
		return
	}

	rspd.Status(http.StatusOK)
	rspd.Response(s.assembler.UserToResponse(user))
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	request := new(types.UpdateUserRequest)
	id, err := getUserIDFromURL(r)
	if err != nil {
		rspd.InternalError(err)
		return
	}

	if err := render.Bind(r, request); err != nil {
		switch err {
		case types.ErrBadRequest:
			rspd.StatusOnly(http.StatusBadRequest)
		default:
			rspd.InternalError(err)
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
		rspd.StatusOnly(http.StatusNotFound)
		return
	default:
		rspd.InternalError(err)
		return
	}

	rspd.Status(http.StatusNoContent)
	rspd.LocationHeader(fmt.Sprintf("/users/%d", id))
	rspd.Response(nil)
}

func (s *Server) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	id, err := getUserIDFromURL(r)
	if err != nil {
		rspd.InternalError(err)
		return
	}

	balance, err := s.app.Queries.GetUserBalance.Handle(id)
	switch err.(type) {
	case nil:
	case domain.ErrUserNotFound:
		rspd.StatusOnly(http.StatusNotFound)
		return
	default:
		rspd.InternalError(err)
		return
	}

	rspd.Status(http.StatusOK)
	rspd.Response(balance.ToFloat())
}

func (s *Server) ChangeUSDBalance(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	id, err := getUserIDFromURL(r)
	if err != nil {
		rspd.InternalError(err)
		return
	}

	request := new(types.ChangeUSDBalanceRequest)
	if err := render.Bind(r, request); err != nil {
		switch err {
		case types.ErrBadRequest:
			rspd.StatusOnly(http.StatusBadRequest)
		default:
			rspd.InternalError(err)
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
		rspd.StatusOnly(http.StatusBadRequest)
		return
	case domain.ErrUserNotFound:
		rspd.StatusOnly(http.StatusNotFound)
		return
	default:
		rspd.StatusOnly(http.StatusInternalServerError)
		return
	}

	rspd.Status(http.StatusNoContent)
	rspd.LocationHeader(fmt.Sprintf("/users/%d", id))
	rspd.Response(nil)
}

func (s *Server) ChangeBTCBalance(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	id, err := getUserIDFromURL(r)
	if err != nil {
		rspd.InternalError(err)
		return
	}

	request := new(types.ChangeBTCBalanceRequest)
	if err := render.Bind(r, request); err != nil {
		switch err {
		case types.ErrBadRequest:
			rspd.StatusOnly(http.StatusBadRequest)
		default:
			rspd.InternalError(err)
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
		rspd.StatusOnly(http.StatusBadRequest)
		return
	case domain.ErrUserNotFound:
		rspd.StatusOnly(http.StatusNotFound)
		return
	default:
		rspd.StatusOnly(http.StatusInternalServerError)
		return
	}

	rspd.Status(http.StatusNoContent)
	rspd.LocationHeader(fmt.Sprintf("/users/%d", id))
	rspd.Response(nil)
}

func (s *Server) GetBTC(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	btc := s.app.Queries.GetBTC.Handle()

	rspd.Status(http.StatusOK)
	rspd.Response(s.assembler.BTCToResponse(btc))
}

func (s *Server) SetBTCPrice(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	request := new(types.SetBTCPriceRequest)
	if err := render.Bind(r, request); err != nil {
		switch err {
		case types.ErrBadRequest:
			rspd.StatusOnly(http.StatusBadRequest)
		default:
			rspd.InternalError(err)
		}
		return
	}

	if err := s.app.Commands.SetBTCPrice.Handle(commands.SetBTCPriceCommand{
		Price: request.Price,
	}); err != nil {
		rspd.InternalError(err)
		return
	}

	rspd.Status(http.StatusNoContent)
	rspd.LocationHeader("/bitcoin")
	rspd.Response(nil)
}
