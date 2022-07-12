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
		name                  string
		cmd                   ChangeBitcoinBalanceCommand
		updateUserFunc        func(uint64, func(*userEntity.User) (*userEntity.User, error)) error
		getPriceFunc          func() bitcoinEntity.BTCPrice
		updateUserCallsAmount int
		getPriceCallsAmount   int
		err                   error
	}{
		{
			name: "invalid action",
			cmd: ChangeBitcoinBalanceCommand{
				Action: "invalid",
			},
			updateUserFunc: func(_ uint64, _ func(*userEntity.User) (*userEntity.User, error)) error {
				return userRepositories.ErrUserNotFound
			},
			err: common.NewServiceError(
				http.StatusBadRequest,
				"Invalid action: invalid",
			),
		},
		{
			name: "user not found",
			cmd: ChangeBitcoinBalanceCommand{
				UserID: 42,
				Action: "buy",
			},
			updateUserFunc: func(_ uint64, _ func(*userEntity.User) (*userEntity.User, error)) error {
				return userRepositories.ErrUserNotFound
			},
			updateUserCallsAmount: 1,
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
			updateUserFunc: func(userID uint64, updateFunc func(*userEntity.User) (*userEntity.User, error)) error {
				_, err := updateFunc(&userEntity.User{})
				return err
			},
			updateUserCallsAmount: 1,
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
			updateUserFunc: func(userID uint64, updateFunc func(*userEntity.User) (*userEntity.User, error)) error {
				zeroUSD, _ := bitcoinEntity.NewUSD(0)
				zeroBTC, _ := bitcoinEntity.NewBTC(0)

				_, err := updateFunc(
					&userEntity.User{
						Balance: userEntity.Balance{
							USD: zeroUSD,
							BTC: zeroBTC,
						},
					},
				)
				return err
			},
			getPriceFunc: func() bitcoinEntity.BTCPrice {
				price, _ := bitcoinEntity.NewUSD(1)
				return bitcoinEntity.NewBTCPrice(price, time.Now())
			},
			updateUserCallsAmount: 1,
			getPriceCallsAmount:   1,
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
			updateUserFunc: func(userID uint64, updateFunc func(*userEntity.User) (*userEntity.User, error)) error {
				zeroUSD, _ := bitcoinEntity.NewUSD(0)
				zeroBTC, _ := bitcoinEntity.NewBTC(0)

				_, err := updateFunc(
					&userEntity.User{
						Balance: userEntity.Balance{
							USD: zeroUSD,
							BTC: zeroBTC,
						},
					},
				)
				return err
			},
			getPriceFunc: func() bitcoinEntity.BTCPrice {
				price, _ := bitcoinEntity.NewUSD(1)
				return bitcoinEntity.NewBTCPrice(price, time.Now())
			},
			updateUserCallsAmount: 1,
			getPriceCallsAmount:   1,
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
			updateUserFunc: func(userID uint64, updateFunc func(*userEntity.User) (*userEntity.User, error)) error {
				userBalance, _ := bitcoinEntity.NewUSD(10)
				zeroBTC, _ := bitcoinEntity.NewBTC(0)

				_, err := updateFunc(
					&userEntity.User{
						Balance: userEntity.Balance{
							USD: userBalance,
							BTC: zeroBTC,
						},
					},
				)
				return err
			},
			getPriceFunc: func() bitcoinEntity.BTCPrice {
				price, _ := bitcoinEntity.NewUSD(1)
				return bitcoinEntity.NewBTCPrice(price, time.Now())
			},
			updateUserCallsAmount: 1,
			getPriceCallsAmount:   1,
		},
		{
			name: "user has enough funds to sell btc",
			cmd: ChangeBitcoinBalanceCommand{
				UserID: 42,
				Action: "sell",
				Amount: 1,
			},
			updateUserFunc: func(userID uint64, updateFunc func(*userEntity.User) (*userEntity.User, error)) error {
				zeroUSD, _ := bitcoinEntity.NewUSD(0)
				userBalance, _ := bitcoinEntity.NewBTC(1)

				_, err := updateFunc(
					&userEntity.User{
						Balance: userEntity.Balance{
							USD: zeroUSD,
							BTC: userBalance,
						},
					},
				)
				return err
			},
			getPriceFunc: func() bitcoinEntity.BTCPrice {
				price, _ := bitcoinEntity.NewUSD(1)
				return bitcoinEntity.NewBTCPrice(price, time.Now())
			},
			updateUserCallsAmount: 1,
			getPriceCallsAmount:   1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &userRepositories.MockUserRepository{UpdateFunc: tc.updateUserFunc}
			bitcoinRepository := &bitcoinRepositories.MockBTCRepository{GetPriceFunc: tc.getPriceFunc}

			service := NewUserService(userRepository, bitcoinRepository)

			err := service.ChangeBitcoinBalance(tc.cmd)

			assert.Equal(t, tc.err, err)
			assert.Len(t, userRepository.UpdateCalls(), tc.updateUserCallsAmount)
			assert.Len(t, bitcoinRepository.GetPriceCalls(), tc.getPriceCallsAmount)
		})
	}
}
