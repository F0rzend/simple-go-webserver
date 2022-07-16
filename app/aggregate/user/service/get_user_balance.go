package service

import (
	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func (us *UserServiceImpl) GetUserBalance(userID uint64) (bitcoinEntity.USD, error) {
	user, err := us.userRepository.Get(userID)
	if err != nil {
		return bitcoinEntity.USD{}, err
	}
	return user.Balance.Total(us.bitcoinRepository.GetPrice()), nil
}
