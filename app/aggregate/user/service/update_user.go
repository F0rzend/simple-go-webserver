package service

import (
	"net/mail"
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type UpdateUser struct {
	ID    uint64
	Name  *string
	Email *string
}

type UpdateUserHandler struct {
	userRepository entity.UserRepository
}

func NewUpdateUserHandler(userRepository entity.UserRepository) (UpdateUserHandler, error) {
	if userRepository == nil {
		return UpdateUserHandler{}, ErrNilUserRepository
	}

	return UpdateUserHandler{
		userRepository: userRepository,
	}, nil
}

func MustNewUpdateUserHandler(userRepository entity.UserRepository) UpdateUserHandler {
	handler, err := NewUpdateUserHandler(userRepository)
	if err != nil {
		panic(err)
	}

	return handler
}

func (h *UpdateUserHandler) Handle(cmd UpdateUser) error {
	return h.userRepository.Update(
		cmd.ID,
		func(user *entity.User) (*entity.User, error) {
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
