package service

import (
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

type UpdateUserCommand struct {
	ID    uint64
	Name  *string
	Email *string
}

func (us *UserServiceImpl) UpdateUser(cmd UpdateUserCommand) error {
	switch err := us.userRepository.Update(
		cmd.ID,
		func(user *entity.User) (*entity.User, error) {
			if cmd.Name != nil {
				user.Name = *cmd.Name
			}

			if cmd.Email != nil {
				addr, err := mail.ParseAddress(*cmd.Email)
				if err != nil {
					return nil, entity.ErrInvalidEmail
				}
				user.Email = addr
			}

			user.UpdatedAt = time.Now()

			return user, nil
		},
	); err {
	case repositories.ErrUserNotFound:
		return common.NewServiceError(
			http.StatusNotFound,
			fmt.Sprintf(
				"User with id %d not found",
				cmd.ID,
			),
		)
	case entity.ErrInvalidEmail:
		return common.NewServiceError(
			http.StatusBadRequest,
			"You must provide a valid email",
		)
	default:
		return err
	}
}
