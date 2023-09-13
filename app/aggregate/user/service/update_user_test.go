package userservice

import (
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/tests"

	"github.com/F0rzend/simple-go-webserver/app/common"

	"github.com/stretchr/testify/assert"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func TestUserService_UpdateUser(t *testing.T) {
	t.Parallel()

	getUserFunc := func(id uint64) (*userentity.User, error) {
		switch id {
		case 1:
			return &userentity.User{}, nil
		default:
			return nil, common.NewFlaggedError("user not found", common.FlagNotFound)
		}
	}
	saveUserFunc := func(user *userentity.User) error {
		return nil
	}

	type command struct {
		userID uint64
		name   *string
		email  *string
	}

	testCases := []struct {
		name                string
		cmd                 command
		getUserCallsAmount  int
		saveUserCallsAmount int
		checkErr            tests.ErrorChecker
	}{
		{
			name: "user not found",
			cmd: command{
				userID: 42,
				name:   nil,
				email:  nil,
			},
			getUserCallsAmount:  1,
			saveUserCallsAmount: 0,
			checkErr:            tests.AssertErrorFlag(common.FlagNotFound),
		},
		{
			name: "update name",
			cmd: command{
				userID: 1,
				name:   strPointer("new name"),
				email:  nil,
			},
			getUserCallsAmount:  1,
			saveUserCallsAmount: 1,
			checkErr:            assert.NoError,
		},
		{
			name: "invalid email",
			cmd: command{
				userID: 1,
				name:   nil,
				email:  strPointer("invalid email"),
			},
			getUserCallsAmount:  1,
			saveUserCallsAmount: 0,
			checkErr:            tests.AssertErrorFlag(common.FlagInvalidArgument),
		},
		{
			name: "update email",
			cmd: command{
				userID: 1,
				name:   nil,
				email:  strPointer("test@mail.com"),
			},
			getUserCallsAmount:  1,
			saveUserCallsAmount: 1,
			checkErr:            assert.NoError,
		},
		{
			name: "update name and email",
			cmd: command{
				userID: 1,
				name:   strPointer("new name"),
				email:  strPointer("test@mail.com"),
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

			err := service.UpdateUser(
				tc.cmd.userID,
				tc.cmd.name,
				tc.cmd.email,
			)

			tc.checkErr(t, err)
			assert.Len(t, userRepository.GetCalls(), tc.getUserCallsAmount)
			assert.Len(t, userRepository.SaveCalls(), tc.saveUserCallsAmount)
		})
	}
}

func strPointer(s string) *string {
	return &s
}
