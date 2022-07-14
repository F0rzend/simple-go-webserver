package service

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	bitcoinRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"

	"github.com/F0rzend/simple-go-webserver/app/common"

	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
)

func TestUserService_GetUser(t *testing.T) {
	t.Parallel()

	getUserFunc := func(userID uint64) (*userEntity.User, error) {
		switch userID {
		case 1:
			return &userEntity.User{}, nil
		default:
			return nil, userRepositories.ErrUserNotFound
		}
	}

	testCases := []struct {
		name   string
		userID uint64
		err    error
	}{
		{
			name:   "success",
			userID: 1,
		},
		{
			name:   "user not found",
			userID: 0,
			err: common.NewServiceError(
				http.StatusNotFound,
				"User with id 0 not found",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &userRepositories.MockUserRepository{GetFunc: getUserFunc}
			bitcoinRepository := &bitcoinRepositories.MockBTCRepository{}

			service := NewUserService(userRepository, bitcoinRepository)

			_, err := service.GetUser(tc.userID)

			assert.Equal(t, tc.err, err)
			assert.Len(t, userRepository.GetCalls(), 1)
		})
	}
}
