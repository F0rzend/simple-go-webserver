package service

import (
	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type UserService struct {
	userRepository    userEntity.UserRepository
	bitcoinRepository bitcoinEntity.BTCRepository

	userIDGenerator func() uint64
}

func NewUserService(
	userRepository userEntity.UserRepository,
	bitcoinRepository bitcoinEntity.BTCRepository,
) *UserService {
	return &UserService{
		userRepository:    userRepository,
		bitcoinRepository: bitcoinRepository,

		userIDGenerator: getUserIDGenerator(),
	}
}
