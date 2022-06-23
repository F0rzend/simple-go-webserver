package commands

import (
	"time"

	"github.com/F0rzend/simple-go-webserver/app/domain"
)

type CreateUserCommand struct {
	Name     string
	Username string
	Email    string
}

type CreateUserCommandHandler struct {
	getID          func() uint64
	userRepository domain.UserRepository
}

func NewCreateUserCommand(userRepository domain.UserRepository) (CreateUserCommandHandler, error) {
	if userRepository == nil {
		return CreateUserCommandHandler{}, ErrNilUserRepository
	}

	return CreateUserCommandHandler{
		getID:          userIDGenerator(),
		userRepository: userRepository,
	}, nil
}

func MustNewCreateUserCommand(userRepository domain.UserRepository) CreateUserCommandHandler {
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

func (h *CreateUserCommandHandler) Handle(cmd CreateUserCommand) (uint64, error) {
	user, err := domain.NewUser(
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
