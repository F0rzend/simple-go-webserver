package service

import (
	"net/http"
	"time"

	userRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"

	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

type CreateUser struct {
	Name     string
	Username string
	Email    string
}

type CreateUserHandler struct {
	getID          func() uint64
	userRepository userEntity.UserRepository
}

func NewCreateUserCommand(userRepository userEntity.UserRepository) (CreateUserHandler, error) {
	if userRepository == nil {
		return CreateUserHandler{}, ErrNilUserRepository
	}

	return CreateUserHandler{
		getID:          userIDGenerator(),
		userRepository: userRepository,
	}, nil
}

func MustNewCreateUserHandler(userRepository userEntity.UserRepository) CreateUserHandler {
	handler, err := NewCreateUserCommand(userRepository)
	if err != nil {
		panic(err)
	}

	return handler
}

func userIDGenerator() func() uint64 {
	var id uint64
	return func() uint64 {
		id++
		return id
	}
}

func (h *CreateUserHandler) Handle(cmd CreateUser) (uint64, error) {
	user, err := userEntity.NewUser(
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
	case userEntity.ErrNameEmpty:
		return 0, common.NewServiceError(
			http.StatusBadRequest,
			"Name cannot be empty",
		)
	case userEntity.ErrUsernameEmpty:
		return 0, common.NewServiceError(
			http.StatusBadRequest,
			"Username cannot be empty",
		)
	case userEntity.ErrInvalidEmail:
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
	case userRepositories.ErrUserAlreadyExists:
		return 0, common.NewServiceError(
			http.StatusBadRequest,
			"This email is already registered",
		)
	default:
		return 0, err
	}
}
