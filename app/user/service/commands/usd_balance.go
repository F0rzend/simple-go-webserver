package commands

import (
	"github.com/F0rzend/simple-go-webserver/app/bitcoin/domain"
	domain2 "github.com/F0rzend/simple-go-webserver/app/user/domain"
)

type ChangeUSDBalanceCommand struct {
	UserID uint64
	Action string
	Amount float64
}

type ChangeUSDBalanceCommandHandler struct {
	userRepository domain2.UserRepository
}

func NewChangeUSDBalanceCommand(
	userRepository domain2.UserRepository,
) (ChangeUSDBalanceCommandHandler, error) {
	if userRepository == nil {
		return ChangeUSDBalanceCommandHandler{}, ErrNilUserRepository
	}

	return ChangeUSDBalanceCommandHandler{
		userRepository: userRepository,
	}, nil
}

func MustNewChangeUSDBalanceCommand(
	userRepository domain2.UserRepository,
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

	return h.userRepository.Update(cmd.UserID, func(user *domain2.User) (*domain2.User, error) {
		usd, err := domain.NewUSD(cmd.Amount)
		if err != nil {
			return nil, err
		}

		if err := user.ChangeUSDBalance(action, usd); err != nil {
			return nil, err
		}
		return user, nil
	})
}
