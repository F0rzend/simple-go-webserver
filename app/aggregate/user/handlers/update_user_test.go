package userhandlers

import (
	"net/http"
	"testing"
	"time"

	userentity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"

	"github.com/stretchr/testify/assert"

	bitcoinservice "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	userservice "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"

	"github.com/F0rzend/simple-go-webserver/app/tests"
)

func TestUserHTTPHandlers_UpdateUser(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                string
		request             UpdateUserRequest
		locationHeader      string
		saveUserCallsAmount int
		expectedStatus      int
	}{
		{
			name:                "empty request",
			request:             UpdateUserRequest{},
			locationHeader:      "",
			saveUserCallsAmount: 0,
			expectedStatus:      http.StatusBadRequest,
		},
		{
			name: "update name and email",
			request: UpdateUserRequest{
				Name:  strPointer("test"),
				Email: strPointer("test@mail.com"),
			},
			locationHeader:      "/users/1",
			saveUserCallsAmount: 1,
			expectedStatus:      http.StatusNoContent,
		},
		{
			name: "correct email",
			request: UpdateUserRequest{
				Email: strPointer("test@mail.com"),
			},
			locationHeader:      "/users/1",
			saveUserCallsAmount: 1,
			expectedStatus:      http.StatusNoContent,
		},
		{
			name: "incorrect email",
			request: UpdateUserRequest{
				Email: strPointer("test"),
			},
			locationHeader:      "",
			saveUserCallsAmount: 0,
			expectedStatus:      http.StatusBadRequest,
		},
		{
			name: "update name",
			request: UpdateUserRequest{
				Name: strPointer("test"),
			},
			locationHeader:      "/users/1",
			saveUserCallsAmount: 1,
			expectedStatus:      http.StatusNoContent,
		},
	}

	getUserIDFromURL := func(_ *http.Request) (uint64, error) {
		return 1, nil
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &userservice.MockUserRepository{
				GetFunc: func(id uint64) (*userentity.User, error) {
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
				},
				SaveFunc: func(_ *userentity.User) error {
					return nil
				},
			}
			bitcoinRepository := &bitcoinservice.MockBTCRepository{}

			service := userservice.NewUserService(userRepository, bitcoinRepository)

			sut := NewUserHTTPHandlers(service, getUserIDFromURL).UpdateUser

			tests.HTTPExpect(t, sut).
				POST("/users/1").
				WithJSON(tc.request).
				Expect().
				Status(tc.expectedStatus).
				ContentType("application/json", "utf-8").
				Header("Location").Equal(tc.locationHeader)

			assert.Len(t, userRepository.SaveCalls(), tc.saveUserCallsAmount)
		})
	}
}

func strPointer(s string) *string {
	return &s
}
