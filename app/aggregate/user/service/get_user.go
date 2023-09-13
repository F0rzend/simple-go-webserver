package userservice

import (
	"fmt"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func (us *UserService) GetUser(userID uint64) (*userentity.User, error) {
	user, err := us.userRepository.Get(userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}
