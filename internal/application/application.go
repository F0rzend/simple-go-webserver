package application

import (
	"github.com/F0rzend/SimpleGoWebserver/internal/application/commands"
	"github.com/F0rzend/SimpleGoWebserver/internal/application/queries"
	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
)

type Application struct {
	Commands commands.Commands
	Queries  queries.Queries
}

func NewApplication(
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
