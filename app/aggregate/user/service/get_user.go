package userservice

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func (us *UserService) GetUser(userID uint64) (*userentity.User, error) {
	return us.userRepository.Get(userID)
}
