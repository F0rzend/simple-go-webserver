package service

import (
	"time"

	"github.com/F0rzend/simple-go-webserver/app/user/repositories"

	"github.com/F0rzend/simple-go-webserver/app/user/service/commands"
	"github.com/F0rzend/simple-go-webserver/app/user/service/queries"

	bitcoinDomain "github.com/F0rzend/simple-go-webserver/app/bitcoin/domain"
	userDomain "github.com/F0rzend/simple-go-webserver/app/user/domain"

	bitcoinRepositories "github.com/F0rzend/simple-go-webserver/app/bitcoin/repositories"
)

type UserService struct {
	Commands commands.Commands
	Queries  queries.Queries
}

func newUserService(
	userRepository userDomain.UserRepository,
	btcRepository bitcoinDomain.BTCRepository,
) *UserService {
	return &UserService{
		Commands: commands.Commands{
			CreateUser:       commands.MustNewCreateUserCommand(userRepository),
			UpdateUser:       commands.MustNewUpdateUserCommand(userRepository),
			ChangeUSDBalance: commands.MustNewChangeUSDBalanceCommand(userRepository),
			ChangeBTCBalance: commands.MustNewChangeBTCBalanceCommand(userRepository, btcRepository),
		},
		Queries: queries.Queries{
			GetUser:        queries.MustNewGetUserQuery(userRepository),
			GetUserBalance: queries.MustNewGetUserBalanceQuery(userRepository, btcRepository),
		},
	}
}

func MustUserService() *UserService {
	userRepository := repositories.NewMemoryUserRepository()
	btcRepository, err := bitcoinRepositories.NewMemoryBTCRepository(bitcoinDomain.MustNewUSD(100))
	if err != nil {
		panic(err)
	}

	return newUserService(userRepository, btcRepository)
}

func NewComponentTestUserService() (*UserService, error) {
	now := time.Now()
	users := map[uint64]*userDomain.User{
		1: userDomain.MustNewUser(
			1,
			"John",
			"Doe",
			"johndoe@mail.com",
			0,
			0,
			now,
			now,
		),
		2: userDomain.MustNewUser(
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
	userRepository := &repositories.MockUserRepository{
		CreateFunc: func(user *userDomain.User) error {
			now := time.Now()
			btc, _ := user.Balance.BTC.ToFloat().Float64()
			usd, _ := user.Balance.USD.ToFloat().Float64()
			_, err := userDomain.NewUser(
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
				return userDomain.ErrUserNotFound(id)
			}
			return nil
		},
		GetFunc: func(id uint64) (*userDomain.User, error) {
			user, ok := users[id]
			if !ok {
				return nil, userDomain.ErrUserNotFound(id)
			}
			return user, nil
		},
		UpdateFunc: func(id uint64, updFunc func(*userDomain.User) (*userDomain.User, error)) error {
			user, ok := users[id]
			if !ok {
				return userDomain.ErrUserNotFound(id)
			}
			userCopy := *user
			_, err := updFunc(&userCopy)
			return err
		},
	}
	btcRepository := &bitcoinRepositories.MockBTCRepository{
		GetFunc: func() bitcoinDomain.BTCPrice {
			return bitcoinDomain.NewBTCPrice(bitcoinDomain.MustNewUSD(100))
		},
		SetPriceFunc: func(price bitcoinDomain.USD) error {
			return nil
		},
	}

	return newUserService(userRepository, btcRepository), nil
}
