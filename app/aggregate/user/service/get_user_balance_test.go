package userservice

import (
	"net/http"
	"testing"

	userrepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/common"
	"github.com/stretchr/testify/assert"
)

func TestUserService_GetUserBalance(t *testing.T) {
	t.Parallel()

	getUserFunc := func(userID uint64) (*userentity.User, error) {
		switch userID {
		case 1:
			return &userentity.User{}, nil
		default:
			return nil, userrepositories.ErrUserNotFound
		}
	}

	getBitcoinPriceFunc := func() bitcoinentity.BTCPrice {
		return bitcoinentity.BTCPrice{}
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

			userRepository := &MockUserRepository{GetFunc: getUserFunc}
			btcPriceGetter := &MockBTCPriceGetter{GetPriceFunc: getBitcoinPriceFunc}

			service := NewUserService(userRepository, btcPriceGetter)

			_, err := service.GetUserBalance(tc.userID)

			assert.Equal(t, tc.err, err)
			assert.Len(t, userRepository.GetCalls(), 1)
			assert.Len(t, btcPriceGetter.GetPriceCalls(), tc.getBitcoinPriceCallsAmount)
		})
	}
}
