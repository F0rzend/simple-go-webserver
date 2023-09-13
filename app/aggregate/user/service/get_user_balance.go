package userservice

import (
	"fmt"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func (us *UserService) GetUserBalance(userID uint64) (bitcoinentity.USD, error) {
	user, err := us.userRepository.Get(userID)
	if err != nil {
		return bitcoinentity.USD{}, fmt.Errorf("error getting user: %w", err)
	}

	currentBTCPrice := us.priceGetter.GetPrice()
	userBalance := user.Balance.Total(currentBTCPrice)

	return userBalance, nil
}
