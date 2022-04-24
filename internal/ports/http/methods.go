package http

import (
	"fmt"
	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"

	"github.com/go-chi/render"

	"github.com/F0rzend/SimpleGoWebserver/internal/application/commands"
	"github.com/F0rzend/SimpleGoWebserver/internal/ports/http/types"
)

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	request := new(types.CreateUserRequest)

	if err := render.Bind(r, request); err != nil {
		var status int
		var response render.Renderer

		switch err.(type) {
		case types.ErrInvalidEmail:
			status = http.StatusBadRequest
			errMsg := fmt.Sprintf("invalid email: %s", err.Error())
			response = Error(http.StatusBadRequest, errMsg)
		default:
			status = http.StatusInternalServerError
			response = types.InternalError
		}
		rspd.Status(status)
		rspd.Response(response)
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

	rspd.Status(http.StatusCreated)
	rspd.LocationHeader(fmt.Sprintf("/users/%d", id))
	rspd.Response(SuccessResponse(fmt.Sprintf("/users/%d", id)))

}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		rspd.InternalError(err)
		return
	}

	user, err := s.app.Queries.GetUser.Handle(id)
	switch err.(type) {
	case nil:
	case domain.ErrUserNotFound:
		rspd.Status(http.StatusNotFound)
		rspd.Response(Error(http.StatusNotFound, fmt.Sprintf("user with id %d not found", id)))
		return
	default:
		rspd.InternalError(err)
		return
	}

	rspd.Status(http.StatusOK)
	rspd.Response(SuccessResponse(s.assembler.UserToResponse(*user)))
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	request := new(types.UpdateUserRequest)
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		rspd.InternalError(err)
		return
	}

	if err := render.Bind(r, request); err != nil {
		var status int
		var response render.Renderer

		switch err.(type) {
		case types.ErrInvalidEmail:
			status = http.StatusBadRequest
			errMsg := err.Error()
			response = Error(http.StatusBadRequest, errMsg)
		default:
			status = http.StatusInternalServerError
			response = types.InternalError
		}
		rspd.Status(status)
		rspd.Response(response)
		return
	}

	if err := s.app.Commands.UpdateUser.Handle(commands.UpdateUserCommand{
		Id:    id,
		Name:  request.Name,
		Email: request.Email,
	}); err != nil {
		rspd.InternalError(err)
		return
	}

	user, err := s.app.Queries.GetUser.Handle(id)
	if err != nil {
		rspd.InternalError(err)
		return
	}

	rspd.Status(http.StatusOK)
	rspd.Response(SuccessResponse(s.assembler.UserToResponse(*user)))

}

func (s *Server) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		rspd.InternalError(err)
		return
	}

	balance, err := s.app.Queries.GetUserBalance.Handle(id)
	switch err.(type) {
	case nil:
	case domain.ErrUserNotFound:
		rspd.Status(http.StatusNotFound)
		rspd.Response(Error(http.StatusNotFound, fmt.Sprintf("user with id %d not found", id)))
		return
	default:
		rspd.InternalError(err)
		return
	}

	rspd.Status(http.StatusOK)
	rspd.Response(SuccessResponse(balance.ToFloat()))

}

func (s *Server) GetBTC(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	btc, err := s.app.Queries.GetBTC.Handle()
	if err != nil {
		rspd.InternalError(err)
		return
	}

	rspd.Status(http.StatusOK)
	rspd.Response(SuccessResponse(s.assembler.BTCToResponse(btc)))
}

func (s *Server) SetBTCPrice(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	request := new(types.SetBTCPriceRequest)
	if err := render.Bind(r, request); err != nil {
		var status int
		var response render.Renderer

		switch err.(type) {
		case types.ErrInvalidPrice:
			status = http.StatusBadRequest
			errMsg := fmt.Sprintf("invalid btc price: %s", err.Error())
			response = Error(http.StatusBadRequest, errMsg)
		default:
			status = http.StatusInternalServerError
			response = types.InternalError
		}
		rspd.Status(status)
		rspd.Response(response)
		return
	}

	if err := s.app.Commands.SetBTCPrice.Handle(commands.SetBTCPriceCommand{
		Price: request.Price,
	}); err != nil {
		rspd.InternalError(err)
		return
	}

	rspd.Status(http.StatusCreated)
	rspd.LocationHeader("/bitcoin")
	rspd.Response(SuccessResponse("/bitcoin"))
}

func (s *Server) ChangeUSDBalance(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		rspd.InternalError(err)
		return
	}

	request := new(types.ChangeUSDBalanceRequest)
	if err := render.Bind(r, request); err != nil {
		var status int
		var response render.Renderer

		switch err {
		case types.ErrInvalidAction:
			status = http.StatusBadRequest
			errMsg := "Invalid action. Action should be 'withdraw' or 'deposit'"
			response = Error(http.StatusBadRequest, errMsg)
		case types.ErrInvalidAmount:
			status = http.StatusBadRequest
			errMsg := "Invalid amount. Amount should be positive number"
			response = Error(http.StatusBadRequest, errMsg)
		default:
			status = http.StatusInternalServerError
			response = types.InternalError
		}
		rspd.Status(status)
		rspd.Response(response)
		return
	}

	err = s.app.Commands.ChangeUSDBalance.Handle(commands.ChangeUSDBalanceCommand{
		UserId: id,
		Action: request.Action,
		Amount: request.Amount,
	})
	switch err.(type) {
	case nil:
	case domain.ErrUserNotFound:
		rspd.Status(http.StatusNotFound)
		rspd.Response(Error(http.StatusNotFound, fmt.Sprintf("user with id %d not found", id)))
		return
	default:
		rspd.InternalError(err)
		return
	}

	rspd.Status(http.StatusOK)
	rspd.LocationHeader(fmt.Sprintf("/users/%d", id))
	rspd.Response(SuccessResponse(fmt.Sprintf("/users/%d", id)))
}

func (s *Server) ChangeBTCBalance(w http.ResponseWriter, r *http.Request) {
	rspd := MustNewResponder(w, r)

	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		rspd.InternalError(err)
		return
	}

	request := new(types.ChangeBTCBalanceRequest)
	if err := render.Bind(r, request); err != nil {
		var status int
		var response render.Renderer

		switch err {
		case types.ErrInvalidAction:
			status = http.StatusBadRequest
			errMsg := "Invalid action. Action should be 'buy' or 'sell'"
			response = Error(http.StatusBadRequest, errMsg)
		case types.ErrInvalidAmount:
			status = http.StatusBadRequest
			errMsg := "Invalid amount. Amount should be positive number"
			response = Error(http.StatusBadRequest, errMsg)
		default:
			status = http.StatusInternalServerError
			response = types.InternalError
		}
		rspd.Status(status)
		rspd.Response(response)
		return
	}

	err = s.app.Commands.ChangeBTCBalance.Handle(commands.ChangeBTCBalanceCommand{
		UserId: id,
		Action: request.Action,
		Amount: request.Amount,
	})
	switch err.(type) {
	case nil:
	case domain.ErrUserNotFound:
		rspd.Status(http.StatusNotFound)
		rspd.Response(Error(http.StatusNotFound, fmt.Sprintf("user with id %d not found", id)))
		return
	default:
		rspd.InternalError(err)
		return
	}

	rspd.Status(http.StatusOK)
	rspd.LocationHeader(fmt.Sprintf("/users/%d", id))
	rspd.Response(SuccessResponse(fmt.Sprintf("/users/%d", id)))
}
