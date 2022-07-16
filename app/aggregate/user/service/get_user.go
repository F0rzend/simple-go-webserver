package service

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func (us *UserServiceImpl) GetUser(userID uint64) (*entity.User, error) {
	return us.userRepository.Get(userID)
}
