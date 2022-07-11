package handlers

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
)

type UserHTTPHandlers struct {
	service service.UserService
}

func NewUserHTTPHandlers(userService service.UserService) *UserHTTPHandlers {
	return &UserHTTPHandlers{
		service: userService,
	}
}
