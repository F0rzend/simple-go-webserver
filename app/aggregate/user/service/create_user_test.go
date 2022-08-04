package userservice

import (
	"net/http"
	"testing"

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
		err             error
	}{
		{
			name: "success",
			cmd: command{
				name:     "test",
				username: "test",
				email:    "test@mail.com",
			},
			saveCallsAmount: 1,
			err:             nil,
		},
		{
			name: "empty name",
			cmd: command{
				name:     "",
				username: "test",
				email:    "test@mail.com",
			},
			err: common.NewApplicationError(
				http.StatusBadRequest,
				"Name cannot be empty",
			),
		},
		{
			name: "empty username",
			cmd: command{
				name:     "test",
				username: "",
				email:    "test@mail.com",
			},
			err: common.NewApplicationError(
				http.StatusBadRequest,
				"Username cannot be empty",
			),
		},
		{
			name: "invalid email",
			cmd: command{
				name:     "test",
				username: "test",
				email:    "test",
			},
			err: common.NewApplicationError(
				http.StatusBadRequest,
				"You must provide a valid email",
			),
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

			assert.Equal(t, tc.err, err)
			assert.Len(t, userRepository.SaveCalls(), tc.saveCallsAmount)
		})
	}
}
