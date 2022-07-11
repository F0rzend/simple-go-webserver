package service

import (
	"fmt"
	"net/http"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

func (us *UserServiceImpl) GetUser(userID uint64) (*entity.User, error) {
	switch user, err := us.userRepository.Get(userID); err {
	case nil:
		return user, nil
	case repositories.ErrUserNotFound:
		return nil, common.NewServiceError(
			http.StatusNotFound,
			fmt.Sprintf(
				"User with id %d not found",
				userID,
			),
		)
	default:
		return nil, err
	}
}
