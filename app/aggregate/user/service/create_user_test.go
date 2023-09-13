package userservice

import (
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/tests"

	"github.com/F0rzend/simple-go-webserver/app/common"

	"github.com/stretchr/testify/assert"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func TestUserService_CreateUser(t *testing.T) {
	t.Parallel()

	type command struct {
		name     string
		username string
		email    string
	}

	testCases := []struct {
		name            string
		cmd             command
		saveCallsAmount int
		checkErr        tests.ErrorChecker
	}{
		{
			name: "success",
			cmd: command{
				name:     "test",
				username: "test",
				email:    "test@mail.com",
			},
			saveCallsAmount: 1,
			checkErr:        assert.NoError,
		},
		{
			name: "empty name",
			cmd: command{
				name:     "",
				username: "test",
				email:    "test@mail.com",
			},
			checkErr: tests.AssertErrorFlag(common.FlagInvalidArgument),
		},
		{
			name: "empty username",
			cmd: command{
				name:     "test",
				username: "",
				email:    "test@mail.com",
			},
			checkErr: tests.AssertErrorFlag(common.FlagInvalidArgument),
		},
		{
			name: "invalid email",
			cmd: command{
				name:     "test",
				username: "test",
				email:    "test",
			},
			checkErr: tests.AssertErrorFlag(common.FlagInvalidArgument),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &MockUserRepository{
				SaveFunc: func(_ *userentity.User) error {
					return nil
				},
			}
			btcPriceGetter := &MockBTCPriceGetter{}

			service := NewUserService(userRepository, btcPriceGetter)

			_, err := service.CreateUser(
				tc.cmd.name,
				tc.cmd.username,
				tc.cmd.email,
			)

			tc.checkErr(t, err)
			assert.Len(t, userRepository.SaveCalls(), tc.saveCallsAmount)
		})
	}
}
