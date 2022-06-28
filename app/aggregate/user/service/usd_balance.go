package service

import (
	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type ChangeUSDBalance struct {
	UserID uint64
	Action string
	Amount float64
}

type ChangeUSDBalanceHandler struct {
	userRepository userEntity.UserRepository
}

func NewChangeUSDBalanceCommand(
	userRepository userEntity.UserRepository,
) (ChangeUSDBalanceHandler, error) {
	if userRepository == nil {
		return ChangeUSDBalanceHandler{}, ErrNilUserRepository
	}

	return ChangeUSDBalanceHandler{
		userRepository: userRepository,
	}, nil
}

func MustNewChangeUSDBalanceHandler(
	userRepository userEntity.UserRepository,
) ChangeUSDBalanceHandler {
	cmd, err := NewChangeUSDBalanceCommand(userRepository)
	if err != nil {
		panic(err)
	}

	return cmd
}

func (h ChangeUSDBalanceHandler) Handle(cmd ChangeUSDBalance) error {
	action, err := bitcoinEntity.NewUSDAction(cmd.Action)
	if err != nil {
		return err
	}

	return h.userRepository.Update(cmd.UserID, func(user *userEntity.User) (*userEntity.User, error) {
		usd, err := bitcoinEntity.NewUSD(cmd.Amount)
		if err != nil {
			return nil, err
		}

		if err := user.ChangeUSDBalance(action, usd); err != nil {
			return nil, err
		}
		return user, nil
	})
}
