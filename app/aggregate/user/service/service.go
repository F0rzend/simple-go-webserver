package service

import (
	"time"

	"github.com/F0rzend/simple-go-webserver/app/common"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	bitcoinRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
)

type UserService struct {
	CreateUser       CreateUserHandler
	UpdateUser       UpdateUserHandler
	ChangeBTCBalance ChangeBTCBalanceHandler
	ChangeUSDBalance ChangeUSDBalanceHandler

	GetUser        GetUserHandler
	GetUserBalance GetUserBalanceHandler
}

func newUserService(
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

func MustUserService() *UserService {
	userRepository := userRepositories.NewMemoryUserRepository()
	btcRepository, err := bitcoinRepositories.NewMemoryBTCRepository(bitcoinEntity.MustNewUSD(100))
	if err != nil {
		panic(err)
	}

	return newUserService(userRepository, btcRepository)
}

func NewComponentTestUserService() (*UserService, error) { // TODO Move to /app/tests
	now := time.Now()
	users := map[uint64]*userEntity.User{
		1: userEntity.MustNewUser(
			1,
			"John",
			"Doe",
			"johndoe@mail.com",
			0,
			0,
			now,
			now,
		),
		2: userEntity.MustNewUser(
			2,
			"Jane",
			"Doe",
			"janedoe@mail.com",
			100,
			100,
			now,
			now,
		),
	}
	userRepository := &userRepositories.MockUserRepository{
		CreateFunc: func(user *userEntity.User) error {
			now := time.Now()
			btc, _ := user.Balance.BTC.ToFloat().Float64()
			usd, _ := user.Balance.USD.ToFloat().Float64()
			_, err := userEntity.NewUser(
				user.ID,
				user.Name,
				user.Username,
				user.Email.Address,
				btc,
				usd,
				now,
				now,
			)
			return err
		},
		DeleteFunc: func(id uint64) error {
			if _, ok := users[id]; !ok {
				return common.ErrUserNotFound(id)
			}
			return nil
		},
		GetFunc: func(id uint64) (*userEntity.User, error) {
			user, ok := users[id]
			if !ok {
				return nil, common.ErrUserNotFound(id)
			}
			return user, nil
		},
		UpdateFunc: func(id uint64, updFunc func(*userEntity.User) (*userEntity.User, error)) error {
			user, ok := users[id]
			if !ok {
				return common.ErrUserNotFound(id)
			}
			userCopy := *user
			_, err := updFunc(&userCopy)
			return err
		},
	}
	btcRepository := &bitcoinRepositories.MockBTCRepository{
		GetFunc: func() bitcoinEntity.BTCPrice {
			return bitcoinEntity.NewBTCPrice(bitcoinEntity.MustNewUSD(100), time.Now())
		},
		SetPriceFunc: func(price bitcoinEntity.USD) error {
			return nil
		},
	}

	return newUserService(userRepository, btcRepository), nil
}
