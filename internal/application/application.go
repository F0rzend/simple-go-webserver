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
			CreateUser: commands.NewCreateUserCommand(userRepository),
			UpdateUser: commands.NewUpdateUserCommand(userRepository),

			SetBTCPrice: commands.NewSetBTCPriceCommand(btcRepository),
		},
		Queries: queries.Queries{
			GetUser: queries.NewGetUserQuery(userRepository),
			GetBTC:  queries.NewGetBTCCommand(btcRepository),
		},
	}
}
