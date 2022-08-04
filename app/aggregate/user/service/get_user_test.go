package userservice

import (
	"net/http"
	"testing"

	userrepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"

	"github.com/stretchr/testify/assert"

	"github.com/F0rzend/simple-go-webserver/app/common"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
)

func TestUserService_GetUser(t *testing.T) {
	t.Parallel()

	getUserFunc := func(userID uint64) (*userentity.User, error) {
		switch userID {
		case 1:
			return &userentity.User{}, nil
		default:
			return nil, userrepositories.ErrUserNotFound
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
			err: common.NewApplicationError(
				http.StatusNotFound,
				"User not found",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &MockUserRepository{GetFunc: getUserFunc}
			btcPriceGetter := &MockBTCPriceGetter{}

			service := NewUserService(userRepository, btcPriceGetter)

			_, err := service.GetUser(tc.userID)

			assert.Equal(t, tc.err, err)
			assert.Len(t, userRepository.GetCalls(), 1)
		})
	}
}
