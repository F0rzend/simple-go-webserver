package service

import (
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type UpdateUserCommand struct {
	UserID uint64
	Name   *string
	Email  *string
}

func (us *UserServiceImpl) UpdateUser(cmd UpdateUserCommand) error {
	user, err := us.userRepository.Get(cmd.UserID)
	if err != nil {
		return err
	}

	if cmd.Name == nil && cmd.Email == nil {
		return nil
	}

	if cmd.Name != nil {
		user.Name = *cmd.Name
	}

	if cmd.Email != nil {
		newEmail, err := entity.ParseEmail(*cmd.Email)
		if err != nil {
			return err
		}
		user.Email = newEmail
	}

	user.UpdatedAt = time.Now()

	return us.userRepository.Save(user)
}
