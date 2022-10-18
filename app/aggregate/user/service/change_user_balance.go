package userservice

import (
	bitcoinentity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func (us *UserService) ChangeUserBalance(userID uint64, action string, amount float64) error {
	user, err := us.userRepository.Get(userID)
	if err != nil {
		return err
	}

	if err := user.ChangeUSDBalance(
		userentity.Action(action),
		bitcoinentity.NewUSD(amount),
	); err != nil {
		return err
	}

	return us.userRepository.Save(user)
}
