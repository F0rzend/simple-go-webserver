package service

import (
	"net/http"
	"time"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

type CreateUserCommand struct {
	Name     string
	Username string
	Email    string
}

func getUserIDGenerator() func() uint64 {
	var id uint64
	return func() uint64 {
		id++
		return id
	}
}

func (us *UserServiceImpl) CreateUser(cmd CreateUserCommand) (uint64, error) {
	user, err := entity.NewUser(
		us.userIDGenerator(),
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

	if err := us.userRepository.Save(user); err != nil {
		return 0, err
	}
	return user.ID, nil
}
