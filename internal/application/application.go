package application

import (
	"github.com/F0rzend/SimpleGoWebserver/internal/application/commands"
	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
)

type Application struct {
	userRepository domain.UserRepository
	Commands       commands.Commands
}

func NewApplication(userRepository domain.UserRepository) *Application {
	return &Application{
		userRepository: userRepository,
		Commands: commands.Commands{
			CreateUser: commands.NewCreateUserCommand(userRepository),
		},
	}
}
