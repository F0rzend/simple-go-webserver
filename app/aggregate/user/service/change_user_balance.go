package service

import (
	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

type ChangeUserBalanceCommand struct {
	UserID uint64
	Action string
	Amount float64
}

func (us *UserServiceImpl) ChangeUserBalance(cmd ChangeUserBalanceCommand) error {
	action, err := bitcoinEntity.NewUSDAction(cmd.Action)
	if err != nil {
		return err
	}

	usd, err := bitcoinEntity.NewUSD(cmd.Amount)
	if err != nil {
		return err
	}

	user, err := us.userRepository.Get(cmd.UserID)
	if err != nil {
		return err
	}

	if err := user.ChangeUSDBalance(action, usd); err != nil {
		return err
	}

	return us.userRepository.Save(user)
}
