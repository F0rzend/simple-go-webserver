package service

import (
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
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

func NewCreateUserCommand(userRepository entity.UserRepository) (CreateUserHandler, error) {
	if userRepository == nil {
		return CreateUserHandler{}, ErrNilUserRepository
	}

	return CreateUserHandler{
		getID:          userIDGenerator(),
		userRepository: userRepository,
	}, nil
}

func MustNewCreateUserHandler(userRepository entity.UserRepository) CreateUserHandler {
	handler, err := NewCreateUserCommand(userRepository)
	if err != nil {
		panic(err)
	}

	return handler
}

func userIDGenerator() func() uint64 {
	var id uint64 = 0
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
	if err != nil {
		return 0, err
	}

	if err := h.userRepository.Create(user); err != nil {
		return 0, err
	}

	return user.ID, nil
}
