package userhandlers

import (
	"net/http"
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/common"

	"github.com/stretchr/testify/assert"

	bitcoinservice "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	userentity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userservice "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
	"github.com/F0rzend/simple-go-webserver/app/tests"
)

func TestUserHTTPHandlers_CreateUser(t *testing.T) {
	t.Parallel()

	type response struct {
		status   int
		location string
	}

	testCases := []struct {
		name                string
		request             CreateUserRequest
		saveUserCallsAmount int
		response            response
	}{
		{
			name: "success",
			request: CreateUserRequest{
				Name:     "Test",
				Username: "test",
				Email:    "test@mail.com",
			},
			saveUserCallsAmount: 1,
			response: response{
				status:   http.StatusCreated,
				location: "/users/1",
			},
		},
		{
			name: "empty name",
			request: CreateUserRequest{
				Name:     "",
				Username: "test",
				Email:    "test@mail.com",
			},
			saveUserCallsAmount: 0,
			response: response{
				status: http.StatusBadRequest,
			},
		},
		{
			name: "empty username",
			request: CreateUserRequest{
				Name:     "Test",
				Username: "",
				Email:    "test@mail.com",
			},
			saveUserCallsAmount: 0,
			response: response{
				status: http.StatusBadRequest,
			},
		},
		{
			name: "invalid email",
			request: CreateUserRequest{
				Name:     "Test",
				Username: "test",
				Email:    "test",
			},
			saveUserCallsAmount: 0,
			response: response{
				status: http.StatusBadRequest,
			},
		},
	}

	saveUserFunc := func(_ *userentity.User) error {
		return nil
	}
	getUserIDFromURL := func(_ *http.Request) (uint64, error) {
		return 1, nil
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &userservice.MockUserRepository{
				SaveFunc: saveUserFunc,
			}
			bitcoinRepository := &bitcoinservice.MockBTCRepository{}
			service := userservice.NewUserService(userRepository, bitcoinRepository)
			handler := NewUserHTTPHandlers(service, getUserIDFromURL).CreateUser
			sut := common.ErrorHandler(handler)

			tests.HTTPExpect(t, sut).
				POST("/").
				WithJSON(tc.request).
				Expect().
				Status(tc.response.status).
				ContentType("application/json").
				Header("Location").Equal(tc.response.location)

			assert.Len(t, userRepository.SaveCalls(), tc.saveUserCallsAmount)
		})
	}
}
