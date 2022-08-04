package userhandlers

import (
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type UserHTTPHandlers struct {
	service UserService

	getUserIDFromRequest func(r *http.Request) (uint64, error)
}

//go:generate moq -out "mock_user_service.gen.go" . UserService:MockUserService
type UserService interface {
	CreateUser(name, username, email string) (uint64, error)
	GetUser(uint64) (*userentity.User, error)
	UpdateUser(userID uint64, name, email *string) error

	GetUserBalance(userID uint64) (bitcoinentity.USD, error)
	ChangeBitcoinBalance(userID uint64, action string, amount float64) error
	ChangeUserBalance(userID uint64, action string, amount float64) error
}

func NewUserHTTPHandlers(
	userService UserService,
	getUserIDFromRequest func(r *http.Request) (uint64, error),
) *UserHTTPHandlers {
	return &UserHTTPHandlers{
		service:              userService,
		getUserIDFromRequest: getUserIDFromRequest,
	}
}
