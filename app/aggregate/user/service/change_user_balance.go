package service

import (
	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func (us *UserService) ChangeUserBalance(userID uint64, action string, amount float64) error {
	usd, err := bitcoinEntity.NewUSD(amount)
	if err != nil {
		return err
	}

	user, err := us.userRepository.Get(userID)
	if err != nil {
		return err
	}

	if err := user.ChangeUSDBalance(entity.Action(action), usd); err != nil {
		return err
	}

	return us.userRepository.Save(user)
}
