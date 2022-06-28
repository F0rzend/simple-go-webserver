package service

import (
	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type UserService struct {
	CreateUser       CreateUserHandler
	UpdateUser       UpdateUserHandler
	ChangeBTCBalance ChangeBTCBalanceHandler
	ChangeUSDBalance ChangeUSDBalanceHandler

	GetUser        GetUserHandler
	GetUserBalance GetUserBalanceHandler
}

func NewUserService(
	userRepository userEntity.UserRepository,
	btcRepository bitcoinEntity.BTCRepository,
) *UserService {
	return &UserService{
		CreateUser:       MustNewCreateUserHandler(userRepository),
		UpdateUser:       MustNewUpdateUserHandler(userRepository),
		ChangeBTCBalance: MustNewChangeBTCBalanceHandler(userRepository, btcRepository),
		ChangeUSDBalance: MustNewChangeUSDBalanceHandler(userRepository),

		GetUser:        MustNewGetUserHandler(userRepository),
		GetUserBalance: MustNewGetUserBalance(userRepository, btcRepository),
	}
}
