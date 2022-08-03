package userservice

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func (us *UserService) ChangeUserBalance(userID uint64, action string, amount float64) error {
	usd, err := bitcoinentity.NewUSD(amount)
	if err != nil {
		return err
	}

	user, err := us.userRepository.Get(userID)
	if err != nil {
		return err
	}

	if err := user.ChangeUSDBalance(userentity.Action(action), usd); err != nil {
		return err
	}

	return us.userRepository.Save(user)
}
