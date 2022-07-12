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

	testCases := []struct {
		name                string
		cmd                 ChangeBitcoinBalanceCommand
		getUserFunc         func(id uint64) (*userEntity.User, error)
		saveUserFunc        func(user *userEntity.User) error
		getBitcoinPriceFunc func() bitcoinEntity.BTCPrice
		getUserCallsAmount  int
		saveUserCallsAmount int
		getPriceCallsAmount int
		err                 error
	}{
		{
			name: "invalid action",
			cmd: ChangeBitcoinBalanceCommand{
				Action: "invalid",
			},
			saveUserFunc: func(user *userEntity.User) error {
				return nil
			},
			err: common.NewServiceError(
				http.StatusBadRequest,
				"You must pass the correct action. Allowed: buy, sell",
			),
		},
		{
			name: "user not found",
			cmd: ChangeBitcoinBalanceCommand{
				UserID: 42,
				Action: "buy",
			},
			getUserFunc: func(id uint64) (*userEntity.User, error) {
				return nil, userRepositories.ErrUserNotFound
			},
			getUserCallsAmount: 1,
			err: common.NewServiceError(
				http.StatusNotFound,
				"User with id 42 not found",
			),
		},
		{
			name: "negative currency",
			cmd: ChangeBitcoinBalanceCommand{
				UserID: 42,
				Action: "buy",
				Amount: -1,
			},
			err: common.NewServiceError(
				http.StatusBadRequest,
				"The amount of currency cannot be negative. Please pass a number greater than 0",
			),
		},
		{
			name: "user has not enough funds to buy btc",
			cmd: ChangeBitcoinBalanceCommand{
				UserID: 42,
				Action: "buy",
				Amount: 1,
			},
			getUserFunc: func(_ uint64) (*userEntity.User, error) {
				return &userEntity.User{}, nil
			},
			getBitcoinPriceFunc: func() bitcoinEntity.BTCPrice {
				price, _ := bitcoinEntity.NewUSD(1)
				return bitcoinEntity.NewBTCPrice(price, time.Now())
			},
			getPriceCallsAmount: 1,
			err: common.NewServiceError(
				http.StatusBadRequest,
				"The user does not have enough funds to buy BTC",
			),
		},
		{
			name: "user has not enough funds to buy btc",
			cmd: ChangeBitcoinBalanceCommand{
				UserID: 42,
				Action: "sell",
				Amount: 1,
			},
			getUserFunc: func(_ uint64) (*userEntity.User, error) {
				return &userEntity.User{}, nil
			},
			getBitcoinPriceFunc: func() bitcoinEntity.BTCPrice {
				price, _ := bitcoinEntity.NewUSD(1)
				return bitcoinEntity.NewBTCPrice(price, time.Now())
			},
			getPriceCallsAmount: 1,
			err: common.NewServiceError(
				http.StatusBadRequest,
				"The user does not have enough funds to sell BTC",
			),
		},
		{
			name: "user has enough funds to buy btc",
			cmd: ChangeBitcoinBalanceCommand{
				UserID: 42,
				Action: "buy",
				Amount: 1,
			},
			getUserFunc: func(_ uint64) (*userEntity.User, error) {
				dollarBalance, _ := bitcoinEntity.NewUSD(1)
				bitcoinBalance, _ := bitcoinEntity.NewBTC(0)

				return &userEntity.User{
					Balance: userEntity.Balance{
						USD: dollarBalance,
						BTC: bitcoinBalance,
					},
				}, nil
			},
			getBitcoinPriceFunc: func() bitcoinEntity.BTCPrice {
				price, _ := bitcoinEntity.NewUSD(1)
				return bitcoinEntity.NewBTCPrice(price, time.Now())
			},
			getUserCallsAmount:  1,
			getPriceCallsAmount: 1,
		},
		{
			name: "user has enough funds to sell btc",
			cmd: ChangeBitcoinBalanceCommand{
				UserID: 42,
				Action: "sell",
				Amount: 1,
			},
			getUserFunc: func(_ uint64) (*userEntity.User, error) {
				dollarBalance, _ := bitcoinEntity.NewUSD(0)
				bitcoinBalance, _ := bitcoinEntity.NewBTC(1)

				return &userEntity.User{
					Balance: userEntity.Balance{
						USD: dollarBalance,
						BTC: bitcoinBalance,
					},
				}, nil
			},
			getBitcoinPriceFunc: func() bitcoinEntity.BTCPrice {
				price, _ := bitcoinEntity.NewUSD(1)
				return bitcoinEntity.NewBTCPrice(price, time.Now())
			},
			getUserCallsAmount:  1,
			getPriceCallsAmount: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &userRepositories.MockUserRepository{SaveFunc: tc.saveUserFunc, GetFunc: tc.getUserFunc}
			bitcoinRepository := &bitcoinRepositories.MockBTCRepository{GetPriceFunc: tc.getBitcoinPriceFunc}

			service := NewUserService(userRepository, bitcoinRepository)

			err := service.ChangeBitcoinBalance(tc.cmd)

			assert.Equal(t, tc.err, err)
			assert.Len(t, userRepository.SaveCalls(), tc.saveUserCallsAmount)
			assert.Len(t, bitcoinRepository.GetPriceCalls(), tc.getPriceCallsAmount)
		})
	}
}
