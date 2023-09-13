package userservice

import (
	"fmt"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func (us *UserService) ChangeBitcoinBalance(userID uint64, action string, amount float64) error {
	user, err := us.userRepository.Get(userID)
	if err != nil {
		return fmt.Errorf("cannot get user: %w", err)
	}

	currentBitcoinPrice := us.priceGetter.GetPrice()

	err = user.ChangeBTCBalance(
		userentity.Action(action),
		bitcoinentity.NewBTC(amount),
		currentBitcoinPrice,
	)
	if err != nil {
		return fmt.Errorf("cannot change user balance: %w", err)
	}

	err = us.userRepository.Save(user)
	if err != nil {
		return fmt.Errorf("cannot save user: %w", err)
	}

	return nil
}
