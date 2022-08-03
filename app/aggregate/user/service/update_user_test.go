package service

import (
	"net/http"
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/common"

	bitcoinRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
	"github.com/stretchr/testify/assert"

	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
)

func TestUserService_UpdateUser(t *testing.T) {
	t.Parallel()

	getUserFunc := func(id uint64) (*userEntity.User, error) {
		switch id {
		case 1:
			return &userEntity.User{}, nil
		default:
			return nil, userRepositories.ErrUserNotFound
		}
	}
	saveUserFunc := func(user *userEntity.User) error {
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
		err                 error
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
			err: common.NewApplicationError(
				http.StatusNotFound,
				"User not found",
			),
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
			err:                 nil,
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
			err: common.NewApplicationError(
				http.StatusBadRequest,
				"You must provide a valid email",
			),
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
			err:                 nil,
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
			err:                 nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &userRepositories.MockUserRepository{SaveFunc: saveUserFunc, GetFunc: getUserFunc}
			bitcoinRepository := &bitcoinRepositories.MockBTCRepository{}

			service := NewUserService(userRepository, bitcoinRepository)

			err := service.UpdateUser(
				tc.cmd.userID,
				tc.cmd.name,
				tc.cmd.email,
			)

			assert.Equal(t, tc.err, err)
			assert.Len(t, userRepository.GetCalls(), tc.getUserCallsAmount)
			assert.Len(t, userRepository.SaveCalls(), tc.saveUserCallsAmount)
		})
	}
}

func strPointer(s string) *string {
	return &s
}
