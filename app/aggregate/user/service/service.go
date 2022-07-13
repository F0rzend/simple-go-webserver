package service

import (
	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type UserService interface {
	CreateUser(cmd CreateUserCommand) (uint64, error)
	GetUser(uint64) (*userEntity.User, error)
	UpdateUser(cmd UpdateUserCommand) error

	GetUserBalance(userID uint64) (bitcoinEntity.USD, error)
	ChangeBitcoinBalance(cmd ChangeBitcoinBalanceCommand) error
	ChangeUserBalance(cmd ChangeUserBalanceCommand) error
}

type UserServiceImpl struct {
	userRepository    userEntity.UserRepository
	bitcoinRepository bitcoinEntity.BTCRepository

	userIDGenerator func() uint64
}

func NewUserService(
	userRepository userEntity.UserRepository,
	bitcoinRepository bitcoinEntity.BTCRepository,
) UserService {
	return &UserServiceImpl{
		userRepository:    userRepository,
		bitcoinRepository: bitcoinRepository,

		userIDGenerator: getUserIDGenerator(),
	}
}
