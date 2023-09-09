package userhandlers

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	bitcoinservice "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	userentity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userservice "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
	"github.com/F0rzend/simple-go-webserver/app/tests"
)

func TestUserHTTPHandlers_UpdateUser(t *testing.T) {
	t.Parallel()

	type response struct {
		status   int
		location string
	}

	testCases := []struct {
		name                string
		request             UpdateUserRequest
		saveUserCallsAmount int
		response            response
	}{
		{
			name:                "empty request",
			request:             UpdateUserRequest{},
			saveUserCallsAmount: 0,
			response: response{
				status: http.StatusBadRequest,
			},
		},
		{
			name: "update name and email",
			request: UpdateUserRequest{
				Name:  strPointer("test"),
				Email: strPointer("test@mail.com"),
			},
			saveUserCallsAmount: 1,
			response: response{
				status:   http.StatusNoContent,
				location: "/users/1",
			},
		},
		{
			name: "correct email",
			request: UpdateUserRequest{
				Email: strPointer("test@mail.com"),
			},
			saveUserCallsAmount: 1,
			response: response{
				status:   http.StatusNoContent,
				location: "/users/1",
			},
		},
		{
			name: "incorrect email",
			request: UpdateUserRequest{
				Email: strPointer("test"),
			},
			saveUserCallsAmount: 0,
			response: response{
				status: http.StatusBadRequest,
			},
		},
		{
			name: "update name",
			request: UpdateUserRequest{
				Name: strPointer("test"),
			},
			saveUserCallsAmount: 1,
			response: response{
				status:   http.StatusNoContent,
				location: "/users/1",
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
			100,
			time.Now(),
			time.Now(),
		)
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
				GetFunc:  getUserFunc,
				SaveFunc: saveUserFunc,
			}
			bitcoinRepository := &bitcoinservice.MockBTCRepository{}
			service := userservice.NewUserService(userRepository, bitcoinRepository)
			sut := NewUserHTTPHandlers(service, getUserIDFromURL).UpdateUser

			tests.HTTPExpect(t, sut).
				POST("/").
				WithJSON(tc.request).
				Expect().
				Status(tc.response.status).
				ContentType("application/json", "utf-8").
				Header("Location").Equal(tc.response.location)

			assert.Len(t, userRepository.SaveCalls(), tc.saveUserCallsAmount)
		})
	}
}

func strPointer(s string) *string {
	return &s
}
