package commands

import (
	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
)

type ChangeUSDBalanceCommand struct {
	UserId uint64
	Action string
	Amount float64
}

type ChangeUSDBalanceCommandHandler struct {
	userRepository domain.UserRepository
}

func NewChangeUSDBalanceCommand(
	userRepository domain.UserRepository,
) (ChangeUSDBalanceCommandHandler, error) {
	if userRepository == nil {
		return ChangeUSDBalanceCommandHandler{}, ErrNilUserRepository
	}

	return ChangeUSDBalanceCommandHandler{
		userRepository: userRepository,
	}, nil
}

func MustNewChangeUSDBalanceCommand(
	userRepository domain.UserRepository,
) ChangeUSDBalanceCommandHandler {
	cmd, err := NewChangeUSDBalanceCommand(userRepository)
	if err != nil {
		panic(err)
	}

	return cmd
}

func (h ChangeUSDBalanceCommandHandler) Handle(cmd ChangeUSDBalanceCommand) error {
	action, err := domain.NewUSDAction(cmd.Action)
	if err != nil {
		return err
	}

	return h.userRepository.Update(cmd.UserId, func(user *domain.User) (*domain.User, error) {
		if err := user.ChangeUSDBalance(action, domain.USDFromFloat(cmd.Amount)); err != nil {
			return nil, err
		}
		return user, nil
	})
}
