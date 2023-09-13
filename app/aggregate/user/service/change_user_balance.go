package userservice

import (
	"fmt"

	bitcoinentity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func (us *UserService) ChangeUserBalance(userID uint64, action string, amount float64) error {
	user, err := us.userRepository.Get(userID)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}

	err = user.ChangeUSDBalance(
		userentity.Action(action),
		bitcoinentity.NewUSD(amount),
	)
	if err != nil {
		return fmt.Errorf("error changing user balance: %w", err)
	}

	err = us.userRepository.Save(user)
	if err != nil {
		return fmt.Errorf("error saving user: %w", err)
	}

	return nil
}
