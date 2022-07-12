package service

import (
	"fmt"
	"net/http"
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
	user, err := us.userRepository.Get(cmd.ID)
	switch err {
	case nil:
	case repositories.ErrUserNotFound:
		return common.NewServiceError(
			http.StatusNotFound,
			fmt.Sprintf(
				"User with id %d not found",
				cmd.ID,
			),
		)
	default:
		return err
	}

	if cmd.Name != nil {
		user.Name = *cmd.Name
	}

	if cmd.Email != nil {
		newEmail, err := entity.ParseEmail(*cmd.Email)
		if err != nil {
			return common.NewServiceError(
				http.StatusBadRequest,
				"You must provide a valid email",
			)
		}
		user.Email = newEmail
	}

	user.UpdatedAt = time.Now()

	if err := us.userRepository.Save(user); err != nil {
		return err
	}

	return nil
}
