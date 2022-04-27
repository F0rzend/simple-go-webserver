package commands

import (
	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
)

type ChangeUSDBalanceCommand struct {
	UserID uint64
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

	return h.userRepository.Update(cmd.UserID, func(user *domain.User) (*domain.User, error) {
		usd, err := domain.USDFromFloat(cmd.Amount)
		if err != nil {
			return nil, err
		}

		if err := user.ChangeUSDBalance(action, usd); err != nil {
			return nil, err
		}
		return user, nil
	})
}
