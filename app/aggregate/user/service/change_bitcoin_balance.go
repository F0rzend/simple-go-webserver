package service

import (
	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func (us *UserService) ChangeBitcoinBalance(userID uint64, action string, amount float64) error {
	btc, err := bitcoinEntity.NewBTC(amount)
	if err != nil {
		return err
	}

	user, err := us.userRepository.Get(userID)
	if err != nil {
		return err
	}

	currentBitcoinPrice := us.bitcoinRepository.GetPrice()

	if err := user.ChangeBTCBalance(entity.Action(action), btc, currentBitcoinPrice); err != nil {
		return err
	}
	return us.userRepository.Save(user)
}
