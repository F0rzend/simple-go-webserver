package commands

import (
	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
	"net/mail"
	"time"
)

type UpdateUserCommand struct {
	Id    uint64
	Name  *string
	Email *string
}

type UpdateUserCommandHandler struct {
	userRepository domain.UserRepository
}

func NewUpdateUserCommand(userRepository domain.UserRepository) *UpdateUserCommandHandler {
	return &UpdateUserCommandHandler{
		userRepository: userRepository,
	}
}

func (h *UpdateUserCommandHandler) Handle(cmd UpdateUserCommand) error {
	return h.userRepository.Update(
		cmd.Id,
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
