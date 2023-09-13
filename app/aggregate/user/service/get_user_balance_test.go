package userservice

import (
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/tests"

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
			return nil, common.NewFlaggedError("user not found", common.FlagNotFound)
		}
	}

	getBitcoinPriceFunc := func() bitcoinentity.BTCPrice {
		return bitcoinentity.BTCPrice{}
	}

	testCases := []struct {
		name                       string
		userID                     uint64
		getBitcoinPriceCallsAmount int
		checkErr                   tests.ErrorChecker
	}{
		{
			name:                       "success",
			getBitcoinPriceCallsAmount: 1,
			userID:                     1,
			checkErr:                   assert.NoError,
		},
		{
			name:     "user not found",
			userID:   0,
			checkErr: tests.AssertErrorFlag(common.FlagNotFound),
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

			tc.checkErr(t, err)
			assert.Len(t, userRepository.GetCalls(), 1)
			assert.Len(t, btcPriceGetter.GetPriceCalls(), tc.getBitcoinPriceCallsAmount)
		})
	}
}
