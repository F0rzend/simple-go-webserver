package service

import (
	"fmt"
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

type GetUserHandler struct {
	userRepository entity.UserRepository
}

func MustNewGetUserHandler(userRepository entity.UserRepository) GetUserHandler {
	if userRepository == nil {
		panic(ErrNilUserRepository)
	}

	return GetUserHandler{
		userRepository: userRepository,
	}
}

func (h *GetUserHandler) Handle(userID uint64) (*entity.User, error) {
	switch user, err := h.userRepository.Get(userID); err {
	case nil:
		return user, nil
	case userRepositories.ErrUserNotFound:
		return nil, common.NewServiceError(
			http.StatusNotFound,
			fmt.Sprintf(
				"User with id %d not found",
				userID,
			),
		)
	default:
		return nil, err
	}
}
