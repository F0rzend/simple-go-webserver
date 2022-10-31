package userservice

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/F0rzend/simple-go-webserver/app/tests"

	userrepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"

	"github.com/stretchr/testify/assert"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func TestUserService_ChangeBitcoinBalance(t *testing.T) {
	t.Parallel()

	var (
		zeroDollar  = bitcoinentity.NewUSD(0)
		zeroBitcoin = bitcoinentity.NewBTC(0)

		oneDollar  = bitcoinentity.NewUSD(1)
		oneBitcoin = bitcoinentity.NewBTC(1)
	)

	testUsers := map[uint64]*userentity.User{
		0: {},
		1: {Balance: userentity.Balance{USD: oneDollar, BTC: zeroBitcoin}},
		2: {Balance: userentity.Balance{USD: zeroDollar, BTC: oneBitcoin}},
	}

	getUserFunc := func(id uint64) (*userentity.User, error) {
		user, ok := testUsers[id]
		if !ok {
			return nil, userrepositories.ErrUserNotFound
		}
		return user, nil
	}
	saveUserFunc := func(user *userentity.User) error {
		return nil
	}
	getBitcoinPriceFunc := func() bitcoinentity.BTCPrice {
		price, err := bitcoinentity.NewBTCPrice(bitcoinentity.NewUSD(1), time.Now())
		require.NoError(t, err)

		return price
	}

	type command struct {
		userID uint64
		action string
		amount float64
	}

	testCases := []struct {
		name                string
		cmd                 command
		getUserCallsAmount  int
		saveUserCallsAmount int
		getPriceCallsAmount int
		expectedErrorStatus int
	}{
		{
			name: "invalid action",
			cmd: command{
				action: "invalid",
			},
			getUserCallsAmount:  1,
			getPriceCallsAmount: 1,
			expectedErrorStatus: http.StatusBadRequest,
		},
		{
			name: "user not found",
			cmd: command{
				userID: 42,
				action: "buy",
			},
			getUserCallsAmount:  1,
			expectedErrorStatus: http.StatusNotFound,
		},
		{
			name: "negative currency",
			cmd: command{
				userID: 0,
				action: "buy",
				amount: -1,
			},
			getPriceCallsAmount: 1,
			expectedErrorStatus: http.StatusBadRequest,
		},
		{
			name: "user has not enough funds to buy btc",
			cmd: command{
				userID: 0,
				action: "buy",
				amount: 1,
			},
			getPriceCallsAmount: 1,
			expectedErrorStatus: http.StatusBadRequest,
		},
		{
			name: "user has not enough funds to sell btc",
			cmd: command{
				userID: 0,
				action: "sell",
				amount: 1,
			},
			getPriceCallsAmount: 1,
			expectedErrorStatus: http.StatusBadRequest,
		},
		{
			name: "user has enough funds to buy btc",
			cmd: command{
				userID: 1,
				action: "buy",
				amount: 1,
			},
			getUserCallsAmount:  1,
			saveUserCallsAmount: 1,
			getPriceCallsAmount: 1,
			expectedErrorStatus: 0,
		},
		{
			name: "user has enough funds to sell btc",
			cmd: command{
				userID: 2,
				action: "sell",
				amount: 1,
			},
			getUserCallsAmount:  1,
			saveUserCallsAmount: 1,
			getPriceCallsAmount: 1,
			expectedErrorStatus: 0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &MockUserRepository{SaveFunc: saveUserFunc, GetFunc: getUserFunc}
			btcPriceGetter := &MockBTCPriceGetter{GetPriceFunc: getBitcoinPriceFunc}

			service := NewUserService(userRepository, btcPriceGetter)

			err := service.ChangeBitcoinBalance(
				tc.cmd.userID,
				tc.cmd.action,
				tc.cmd.amount,
			)

			tests.ExpectApplicationError(t, tc.expectedErrorStatus, err)
			assert.Len(t, userRepository.SaveCalls(), tc.saveUserCallsAmount)
			assert.Len(t, btcPriceGetter.GetPriceCalls(), tc.getPriceCallsAmount)
		})
	}
}
