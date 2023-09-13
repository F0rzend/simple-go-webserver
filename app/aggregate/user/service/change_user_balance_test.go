package userservice

import (
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/common"

	"github.com/F0rzend/simple-go-webserver/app/tests"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserService_ChangeUserBalance(t *testing.T) {
	t.Parallel()

	testUsers := map[uint64]userentity.User{
		0: {},
		1: {Balance: userentity.Balance{USD: bitcoinentity.NewUSD(1), BTC: bitcoinentity.NewBTC(0)}},
	}

	getUserFunc := func(id uint64) (*userentity.User, error) {
		user, ok := testUsers[id]
		if !ok {
			return nil, common.NewFlaggedError("user not found", common.FlagNotFound)
		}
		return &user, nil
	}
	saveUserFunc := func(user *userentity.User) error {
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
		checkErr            tests.ErrorChecker
	}{
		{
			name: "invalid action",
			cmd: command{
				userID: 0,
				action: "invalid",
				amount: 1,
			},
			getUserCallsAmount: 1,
			checkErr:           tests.AssertErrorFlag(common.FlagInvalidArgument),
		},
		{
			name: "negative currency",
			cmd: command{
				userID: 0,
				action: "deposit",
				amount: -1,
			},
			getUserCallsAmount:  1,
			saveUserCallsAmount: 0,
			checkErr:            tests.AssertErrorFlag(common.FlagInvalidArgument),
		},
		{
			name: "user not found",
			cmd: command{
				userID: 42,
				action: "deposit",
				amount: 1,
			},
			getUserCallsAmount: 1,
			checkErr:           tests.AssertErrorFlag(common.FlagNotFound),
		},
		{
			name: "user has not enough money to withdraw",
			cmd: command{
				userID: 0,
				action: "withdraw",
				amount: 1,
			},
			getUserCallsAmount: 1,
			checkErr:           tests.AssertErrorFlag(common.FlagInvalidArgument),
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
			checkErr:            assert.NoError,
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
			checkErr:            assert.NoError,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &MockUserRepository{SaveFunc: saveUserFunc, GetFunc: getUserFunc}
			btcPriceGetter := &MockBTCPriceGetter{}

			service := NewUserService(userRepository, btcPriceGetter)

			err := service.ChangeUserBalance(
				tc.cmd.userID,
				tc.cmd.action,
				tc.cmd.amount,
			)

			tc.checkErr(t, err)
			assert.Len(t, userRepository.GetCalls(), tc.getUserCallsAmount)
			assert.Len(t, userRepository.SaveCalls(), tc.saveUserCallsAmount)
		})
	}
}
