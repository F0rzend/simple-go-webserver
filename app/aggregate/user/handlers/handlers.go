package handlers

import (
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
)

type UserHTTPHandlers struct {
	service service.UserService

	getUserIDFromRequest func(r *http.Request) (uint64, error)
}

func NewUserHTTPHandlers(
	userService service.UserService,
	getUserIDFromRequest func(r *http.Request) (uint64, error),
) *UserHTTPHandlers {
	return &UserHTTPHandlers{
		service:              userService,
		getUserIDFromRequest: getUserIDFromRequest,
	}
}
