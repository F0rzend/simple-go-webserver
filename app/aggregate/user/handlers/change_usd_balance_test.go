package userhandlers

import (
	"net/http"
	"testing"
	"time"

	"github.com/F0rzend/simple-go-webserver/app/common"

	"github.com/stretchr/testify/assert"

	bitcoinservice "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	userentity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userservice "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"

	"github.com/F0rzend/simple-go-webserver/app/tests"
)

func TestUserHTTPHandlers_ChangeUSDBalance(t *testing.T) {
	t.Parallel()

	type calls struct {
		getUser  int
		saveUser int
	}

	type response struct {
		status   int
		location string
	}

	testCases := []struct {
		name     string
		request  ChangeUSDBalanceRequest
		calls    calls
		response response
	}{
		{
			name: "success deposit",
			request: ChangeUSDBalanceRequest{
				Action: "deposit",
				Amount: 10.0,
			},
			calls: calls{
				getUser:  1,
				saveUser: 1,
			},
			response: response{
				status:   http.StatusNoContent,
				location: "/users/1",
			},
		},
		{
			name: "success withdraw",
			request: ChangeUSDBalanceRequest{
				Action: "withdraw",
				Amount: 1.0,
			},
			calls: calls{
				getUser:  1,
				saveUser: 1,
			},
			response: response{
				status:   http.StatusNoContent,
				location: "/users/1",
			},
		},
		{
			name: "empty action",
			request: ChangeUSDBalanceRequest{
				Amount: 10.0,
			},
			response: response{
				status: http.StatusBadRequest,
			},
		},
		{
			name: "empty amount",
			request: ChangeUSDBalanceRequest{
				Action: "deposit",
			},
			response: response{
				status: http.StatusBadRequest,
			},
		},
	}

	getUserFunc := func(id uint64) (*userentity.User, error) {
		return userentity.NewUser(
			id,
			"John",
			"john",
			"john@mail.com",
			0,
			1,
			time.Now(),
			time.Now(),
		)
	}
	saveUserFunc := func(_ *userentity.User) error {
		return nil
	}
	idProvider := func(_ *http.Request) (uint64, error) {
		return 1, nil
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &userservice.MockUserRepository{
				GetFunc:  getUserFunc,
				SaveFunc: saveUserFunc,
			}
			bitcoinRepository := &bitcoinservice.MockBTCRepository{}
			service := userservice.NewUserService(userRepository, bitcoinRepository)
			handler := NewUserHTTPHandlers(service, idProvider).ChangeUSDBalance
			sut := common.ErrorHandler(handler)

			tests.HTTPExpect(t, sut).
				POST("/").
				WithJSON(tc.request).
				Expect().
				Status(tc.response.status).
				ContentType("application/json").
				Header("Location").Equal(tc.response.location)

			assert.Len(t, userRepository.GetCalls(), tc.calls.getUser)
			assert.Len(t, userRepository.SaveCalls(), tc.calls.saveUser)
		})
	}
}
