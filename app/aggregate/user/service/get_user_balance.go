package service

import (
	"fmt"
	"net/http"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	userRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

func (us *UserServiceImpl) GetUserBalance(userID uint64) (bitcoinEntity.USD, error) {
	user, err := us.userRepository.Get(userID)
	switch err {
	case nil:
		return user.Balance.Total(us.bitcoinRepository.GetPrice()), nil
	case userRepositories.ErrUserNotFound:
		return bitcoinEntity.USD{}, common.NewServiceError(
			http.StatusNotFound,
			fmt.Sprintf(
				"User with id %d not found",
				userID,
			),
		)
	default:
		return bitcoinEntity.USD{}, err
	}
}
