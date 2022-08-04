package userservice

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func (us *UserService) GetUserBalance(userID uint64) (bitcoinentity.USD, error) {
	user, err := us.userRepository.Get(userID)
	if err != nil {
		return bitcoinentity.USD{}, err
	}
	return user.Balance.Total(us.priceGetter.GetPrice()), nil
}
