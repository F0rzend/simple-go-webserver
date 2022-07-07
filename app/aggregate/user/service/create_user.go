package service

import (
	"net/http"
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

type CreateUser struct {
	Name     string
	Username string
	Email    string
}

type CreateUserHandler struct {
	getID          func() uint64
	userRepository entity.UserRepository
}

func MustNewCreateUserHandler(userRepository entity.UserRepository) CreateUserHandler {
	if userRepository == nil {
		panic(ErrNilUserRepository)
	}

	return CreateUserHandler{
		getID:          userIDGenerator(),
		userRepository: userRepository,
	}
}

func userIDGenerator() func() uint64 {
	var id uint64
	return func() uint64 {
		id++
		return id
	}
}

func (h *CreateUserHandler) Handle(cmd CreateUser) (uint64, error) {
	user, err := entity.NewUser(
		h.getID(),
		cmd.Name,
		cmd.Username,
		cmd.Email,
		0,
		0,
		time.Now(),
		time.Now(),
	)
	switch err {
	case nil:
	case entity.ErrNameEmpty:
		return 0, common.NewServiceError(
			http.StatusBadRequest,
			"Name cannot be empty",
		)
	case entity.ErrUsernameEmpty:
		return 0, common.NewServiceError(
			http.StatusBadRequest,
			"Username cannot be empty",
		)
	case entity.ErrInvalidEmail:
		return 0, common.NewServiceError(
			http.StatusBadRequest,
			"You must provide a valid email",
		)
	default:
		return 0, err
	}

	switch err := h.userRepository.Create(user); err {
	case nil:
		return user.ID, nil
	case repositories.ErrUserAlreadyExists:
		return 0, common.NewServiceError(
			http.StatusBadRequest,
			"This email is already registered",
		)
	default:
		return 0, err
	}
}
