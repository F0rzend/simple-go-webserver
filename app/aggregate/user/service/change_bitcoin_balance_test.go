package service

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	bitcoinRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
	"github.com/F0rzend/simple-go-webserver/app/common"
)

func TestUserService_ChangeBitcoinBalance(t *testing.T) {
	t.Parallel()

	var (
		zeroDollar, _  = bitcoinEntity.NewUSD(0)
		zeroBitcoin, _ = bitcoinEntity.NewBTC(0)

		oneDollar, _  = bitcoinEntity.NewUSD(1)
		oneBitcoin, _ = bitcoinEntity.NewBTC(1)
	)

	testUsers := map[uint64]*userEntity.User{
		0: {},
		1: {Balance: userEntity.Balance{USD: oneDollar, BTC: zeroBitcoin}},
		2: {Balance: userEntity.Balance{USD: zeroDollar, BTC: oneBitcoin}},
	}

	getUserFunc := func(id uint64) (*userEntity.User, error) {
		user, ok := testUsers[id]
		if !ok {
			return nil, userRepositories.ErrUserNotFound
		}
		return user, nil
	}
	saveUserFunc := func(user *userEntity.User) error {
		return nil
	}
	getBitcoinPriceFunc := func() bitcoinEntity.BTCPrice {
		price, _ := bitcoinEntity.NewUSD(1)
		return bitcoinEntity.NewBTCPrice(price, time.Now())
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
		err                 error
	}{
		{
			name: "invalid action",
			cmd: command{
				action: "invalid",
			},
			getUserCallsAmount:  1,
			getPriceCallsAmount: 1,
			err: common.NewApplicationError(
				http.StatusBadRequest,
				"You must specify a valid action. Available actions: buy and sell",
			),
		},
		{
			name: "user not found",
			cmd: command{
				userID: 42,
				action: "buy",
			},
			getUserCallsAmount: 1,
			err: common.NewApplicationError(
				http.StatusNotFound,
				"User not found",
			),
		},
		{
			name: "negative currency",
			cmd: command{
				userID: 0,
				action: "buy",
				amount: -1,
			},
			err: common.NewApplicationError(
				http.StatusBadRequest,
				"The amount of currency cannot be negative. Please pass a number greater than 0",
			),
		},
		{
			name: "user has not enough funds to buy btc",
			cmd: command{
				userID: 0,
				action: "buy",
				amount: 1,
			},
			getPriceCallsAmount: 1,
			err: common.NewApplicationError(
				http.StatusBadRequest,
				"The user does not have enough funds",
			),
		},
		{
			name: "user has not enough funds to sell btc",
			cmd: command{
				userID: 0,
				action: "sell",
				amount: 1,
			},
			getPriceCallsAmount: 1,
			err: common.NewApplicationError(
				http.StatusBadRequest,
				"The user does not have enough funds",
			),
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
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &userRepositories.MockUserRepository{SaveFunc: saveUserFunc, GetFunc: getUserFunc}
			bitcoinRepository := &bitcoinRepositories.MockBTCRepository{GetPriceFunc: getBitcoinPriceFunc}

			service := NewUserService(userRepository, bitcoinRepository)

			err := service.ChangeBitcoinBalance(
				tc.cmd.userID,
				tc.cmd.action,
				tc.cmd.amount,
			)

			assert.Equal(t, tc.err, err)
			assert.Len(t, userRepository.SaveCalls(), tc.saveUserCallsAmount)
			assert.Len(t, bitcoinRepository.GetPriceCalls(), tc.getPriceCallsAmount)
		})
	}
}
