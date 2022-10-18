package userservice

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func (us *UserService) ChangeBitcoinBalance(userID uint64, action string, amount float64) error {
	user, err := us.userRepository.Get(userID)
	if err != nil {
		return err
	}

	currentBitcoinPrice := us.priceGetter.GetPrice()

	if err := user.ChangeBTCBalance(
		userentity.Action(action),
		bitcoinentity.NewBTC(amount),
		currentBitcoinPrice,
	); err != nil {
		return err
	}
	return us.userRepository.Save(user)
}
