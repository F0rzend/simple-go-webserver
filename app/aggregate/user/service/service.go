package userservice

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type UserService struct {
	userRepository    userentity.UserRepository
	bitcoinRepository bitcoinentity.BTCRepository

	userIDGenerator func() uint64
}

func NewUserService(
	userRepository userentity.UserRepository,
	bitcoinRepository bitcoinentity.BTCRepository,
) *UserService {
	return &UserService{
		userRepository:    userRepository,
		bitcoinRepository: bitcoinRepository,

		userIDGenerator: getUserIDGenerator(),
	}
}
