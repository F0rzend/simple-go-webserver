package service

import (
	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

type ChangeBitcoinBalanceCommand struct {
	UserID uint64
	Action string
	Amount float64
}

func (us *UserServiceImpl) ChangeBitcoinBalance(cmd ChangeBitcoinBalanceCommand) error {
	btc, err := bitcoinEntity.NewBTC(cmd.Amount)
	if err != nil {
		return err
	}

	user, err := us.userRepository.Get(cmd.UserID)
	if err != nil {
		return err
	}

	currentBitcoinPrice := us.bitcoinRepository.GetPrice()

	if err := user.ChangeBTCBalance(entity.Action(cmd.Action), btc, currentBitcoinPrice); err != nil {
		return err
	}
	return us.userRepository.Save(user)
}
