package http

import (
	"fmt"
	"net/http"

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
			response = Error(http.StatusBadRequest, err)
		default:
			status = http.StatusInternalServerError
			response = types.InternalError
		}
		rspd.Status(status)
		rspd.Response(response)
		return
	}

	id, err := s.app.Commands.CreateUser.Execute(commands.CreateUserCommand{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
	})
	if err != nil {
		rspd.InternalError()
		return
	}

	rspd.Status(http.StatusCreated)
	rspd.LocationHeader(fmt.Sprintf("/users/%d", id))
	rspd.Response(SuccessResponse(nil))

}
