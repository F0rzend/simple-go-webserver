package commands

import (
	"net/mail"
	"time"

	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
)

type UpdateUserCommand struct {
	ID    uint64
	Name  *string
	Email *string
}

type UpdateUserCommandHandler struct {
	userRepository domain.UserRepository
}

func NewUpdateUserCommand(userRepository domain.UserRepository) (UpdateUserCommandHandler, error) {
	if userRepository == nil {
		return UpdateUserCommandHandler{}, ErrNilUserRepository
	}

	return UpdateUserCommandHandler{
		userRepository: userRepository,
	}, nil
}

func MustNewUpdateUserCommand(userRepository domain.UserRepository) UpdateUserCommandHandler {
	handler, err := NewUpdateUserCommand(userRepository)
	if err != nil {
		panic(err)
	}

	return handler
}

func (h *UpdateUserCommandHandler) Handle(cmd UpdateUserCommand) error {
	return h.userRepository.Update(
		cmd.ID,
		func(user *domain.User) (*domain.User, error) {
			if cmd.Name != nil {
				user.Name = *cmd.Name
			}

			if cmd.Email != nil {
				addr, err := mail.ParseAddress(*cmd.Email)
				if err != nil {
					return nil, err
				}
				user.Email = addr
			}

			user.UpdatedAt = time.Now()

			return user, nil
		},
	)
}
