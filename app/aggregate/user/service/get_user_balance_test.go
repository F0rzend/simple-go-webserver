package service

import (
	"net/http"
	"testing"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"

	bitcoinRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
	"github.com/F0rzend/simple-go-webserver/app/common"
	"github.com/stretchr/testify/assert"
)

func TestUserService_GetUserBalance(t *testing.T) {
	t.Parallel()

	getUserFunc := func(userID uint64) (*userEntity.User, error) {
		switch userID {
		case 1:
			return &userEntity.User{}, nil
		default:
			return nil, userRepositories.ErrUserNotFound
		}
	}

	getBitcoinPriceFunc := func() bitcoinEntity.BTCPrice {
		return bitcoinEntity.BTCPrice{}
	}

	testCases := []struct {
		name                       string
		userID                     uint64
		getBitcoinPriceCallsAmount int
		err                        error
	}{
		{
			name:                       "success",
			getBitcoinPriceCallsAmount: 1,
			userID:                     1,
		},
		{
			name:   "user not found",
			userID: 0,
			err: common.NewApplicationError(
				http.StatusNotFound,
				"User not found",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &userRepositories.MockUserRepository{GetFunc: getUserFunc}
			bitcoinRepository := &bitcoinRepositories.MockBTCRepository{GetPriceFunc: getBitcoinPriceFunc}

			service := NewUserService(userRepository, bitcoinRepository)

			_, err := service.GetUserBalance(tc.userID)

			assert.Equal(t, tc.err, err)
			assert.Len(t, userRepository.GetCalls(), 1)
			assert.Len(t, bitcoinRepository.GetPriceCalls(), tc.getBitcoinPriceCallsAmount)
		})
	}
}
