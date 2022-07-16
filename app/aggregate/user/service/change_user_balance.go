package service

import (
	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type ChangeUserBalanceCommand struct {
	UserID uint64
	Action string
	Amount float64
}

func (us *UserServiceImpl) ChangeUserBalance(cmd ChangeUserBalanceCommand) error {
	usd, err := bitcoinEntity.NewUSD(cmd.Amount)
	if err != nil {
		return err
	}

	user, err := us.userRepository.Get(cmd.UserID)
	if err != nil {
		return err
	}

	if err := user.ChangeUSDBalance(entity.Action(cmd.Action), usd); err != nil {
		return err
	}

	return us.userRepository.Save(user)
}
