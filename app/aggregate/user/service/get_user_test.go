package userservice

import (
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/tests"

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
			return nil, common.NewFlaggedError("user not found", common.FlagNotFound)
		}
	}

	testCases := []struct {
		name     string
		userID   uint64
		checkErr tests.ErrorChecker
	}{
		{
			name:     "success",
			userID:   1,
			checkErr: assert.NoError,
		},
		{
			name:     "user not found",
			userID:   0,
			checkErr: tests.AssertErrorFlag(common.FlagNotFound),
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

			tc.checkErr(t, err)
			assert.Len(t, userRepository.GetCalls(), 1)
		})
	}
}
