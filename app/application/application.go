package application

import (
	"time"

	"github.com/F0rzend/simple-go-webserver/app/adapters/repositories/btc"
	"github.com/F0rzend/simple-go-webserver/app/adapters/repositories/user"
	"github.com/F0rzend/simple-go-webserver/app/application/commands"
	"github.com/F0rzend/simple-go-webserver/app/application/queries"
	"github.com/F0rzend/simple-go-webserver/app/domain"
)

type Application struct {
	Commands commands.Commands
	Queries  queries.Queries
}

var DefaultBitcoinPrice = domain.MustNewUSD(100) // nolint: gomnd

func newApplication(
	userRepository domain.UserRepository,
	btcRepository domain.BTCRepository,
) *Application {
	return &Application{
		Commands: commands.Commands{
			CreateUser:       commands.MustNewCreateUserCommand(userRepository),
			UpdateUser:       commands.MustNewUpdateUserCommand(userRepository),
			ChangeUSDBalance: commands.MustNewChangeUSDBalanceCommand(userRepository),
			ChangeBTCBalance: commands.MustNewChangeBTCBalanceCommand(userRepository, btcRepository),

			SetBTCPrice: commands.MustNewSetBTCPriceCommand(btcRepository),
		},
		Queries: queries.Queries{
			GetUser:        queries.MustNewGetUserQuery(userRepository),
			GetUserBalance: queries.MustNewGetUserBalanceQuery(userRepository, btcRepository),

			GetBTC: queries.MustNewGetBTCCommand(btcRepository),
		},
	}
}

func NewApplication() (*Application, error) {
	userRepository := userrepositories.NewMemoryUserRepository()
	btcRepository, err := btcrepositories.NewMemoryBTCRepository(DefaultBitcoinPrice)
	if err != nil {
		return nil, err
	}

	return newApplication(userRepository, btcRepository), nil
}

func NewComponentTestApplication() (*Application, error) {
	now := time.Now()
	users := map[uint64]*domain.User{
		1: domain.MustNewUser(
			1,
			"John",
			"Doe",
			"johndoe@mail.com",
			0,
			0,
			now,
			now,
		),
		2: domain.MustNewUser(
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
	userRepository := &userrepositories.MockUserRepository{
		CreateFunc: func(user *domain.User) error {
			now := time.Now()
			btc, _ := user.Balance.BTC.ToFloat().Float64()
			usd, _ := user.Balance.USD.ToFloat().Float64()
			_, err := domain.NewUser(
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
				return domain.ErrUserNotFound(id)
			}
			return nil
		},
		GetFunc: func(id uint64) (*domain.User, error) {
			user, ok := users[id]
			if !ok {
				return nil, domain.ErrUserNotFound(id)
			}
			return user, nil
		},
		UpdateFunc: func(id uint64, updFunc func(*domain.User) (*domain.User, error)) error {
			user, ok := users[id]
			if !ok {
				return domain.ErrUserNotFound(id)
			}
			userCopy := *user
			_, err := updFunc(&userCopy)
			return err
		},
	}
	btcRepository := &btcrepositories.MockBTCRepository{
		GetFunc: func() domain.BTCPrice {
			return domain.NewBTCPrice(domain.MustNewUSD(100))
		},
		SetPriceFunc: func(price domain.USD) error {
			return nil
		},
	}

	return newApplication(userRepository, btcRepository), nil
}
