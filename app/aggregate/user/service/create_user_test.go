package service

import (
	"net/http"
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/common"

	"github.com/stretchr/testify/assert"

	bitcoinRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
)

func TestUserService_CreateUser(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name            string
		cmd             CreateUserCommand
		saveCallsAmount int
		err             error
	}{
		{
			name: "success",
			cmd: CreateUserCommand{
				Name:     "test",
				Username: "test",
				Email:    "test@mail.com",
			},
			saveCallsAmount: 1,
			err:             nil,
		},
		{
			name: "empty name",
			cmd: CreateUserCommand{
				Name:     "",
				Username: "test",
				Email:    "test@mail.com",
			},
			err: common.NewServiceError(
				http.StatusBadRequest,
				"Name cannot be empty",
			),
		},
		{
			name: "empty username",
			cmd: CreateUserCommand{
				Name:     "test",
				Username: "",
				Email:    "test@mail.com",
			},
			err: common.NewServiceError(
				http.StatusBadRequest,
				"Username cannot be empty",
			),
		},
		{
			name: "invalid email",
			cmd: CreateUserCommand{
				Name:     "test",
				Username: "test",
				Email:    "test",
			},
			err: common.NewServiceError(
				http.StatusBadRequest,
				"You must provide a valid email",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &userRepositories.MockUserRepository{
				SaveFunc: func(_ *userEntity.User) error {
					return nil
				},
			}
			bitcoinRepository := &bitcoinRepositories.MockBTCRepository{}

			service := NewUserService(userRepository, bitcoinRepository)

			_, err := service.CreateUser(tc.cmd)

			assert.Equal(t, tc.err, err)
			assert.Len(t, userRepository.SaveCalls(), tc.saveCallsAmount)
		})
	}
}
