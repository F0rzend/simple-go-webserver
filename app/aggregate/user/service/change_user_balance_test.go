package service

import (
	"net/http"
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/common"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	bitcoinRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
	"github.com/stretchr/testify/assert"
)

func TestUserService_ChangeUserBalance(t *testing.T) {
	t.Parallel()

	var (
		zeroBitcoin, _ = bitcoinEntity.NewBTC(0)
		oneDollar, _   = bitcoinEntity.NewUSD(1)
	)

	testUsers := map[uint64]userEntity.User{
		0: {},
		1: {Balance: userEntity.Balance{USD: oneDollar, BTC: zeroBitcoin}},
	}

	getUserFunc := func(id uint64) (*userEntity.User, error) {
		user, ok := testUsers[id]
		if !ok {
			return nil, userRepositories.ErrUserNotFound
		}
		return &user, nil
	}
	saveUserFunc := func(user *userEntity.User) error {
		return nil
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
		err                 error
	}{
		{
			name: "invalid action",
			cmd: command{
				userID: 0,
				action: "invalid",
				amount: 1,
			},
			getUserCallsAmount: 1,
			err: common.NewApplicationError(
				http.StatusBadRequest,
				"You must specify a valid action. Available actions: deposit and withdraw",
			),
		},
		{
			name: "negative currency",
			cmd: command{
				userID: 0,
				action: "deposit",
				amount: -1,
			},
			err: common.NewApplicationError(
				http.StatusBadRequest,
				"The amount of currency cannot be negative. Please pass a number greater than 0",
			),
		},
		{
			name: "user not found",
			cmd: command{
				userID: 42,
				action: "deposit",
				amount: 1,
			},
			getUserCallsAmount: 1,
			err: common.NewApplicationError(
				http.StatusNotFound,
				"User not found",
			),
		},
		{
			name: "user has not enough money to withdraw",
			cmd: command{
				userID: 0,
				action: "withdraw",
				amount: 1,
			},
			getUserCallsAmount: 1,
			err: common.NewApplicationError(
				http.StatusBadRequest,
				"The user does not have enough funds",
			),
		},
		{
			name: "success withdraw",
			cmd: command{
				userID: 1,
				action: "withdraw",
				amount: 1,
			},
			getUserCallsAmount:  1,
			saveUserCallsAmount: 1,
		},
		{
			name: "success deposit",
			cmd: command{
				userID: 0,
				action: "deposit",
				amount: 1,
			},
			getUserCallsAmount:  1,
			saveUserCallsAmount: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &userRepositories.MockUserRepository{SaveFunc: saveUserFunc, GetFunc: getUserFunc}
			bitcoinRepository := &bitcoinRepositories.MockBTCRepository{}

			service := NewUserService(userRepository, bitcoinRepository)

			err := service.ChangeUserBalance(
				tc.cmd.userID,
				tc.cmd.action,
				tc.cmd.amount,
			)

			assert.Equal(t, tc.err, err)
			assert.Len(t, userRepository.GetCalls(), tc.getUserCallsAmount)
			assert.Len(t, userRepository.SaveCalls(), tc.saveUserCallsAmount)
		})
	}
}
