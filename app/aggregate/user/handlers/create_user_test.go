package userhandlers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	userentity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"

	bitcoinservice "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	userservice "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"

	"github.com/F0rzend/simple-go-webserver/app/tests"
)

func TestUserHTTPHandlers_CreateUser(t *testing.T) {
	t.Parallel()

	getUserIDFromURL := func(_ *http.Request) (uint64, error) {
		return 1, nil
	}

	testCases := []struct {
		name                string
		request             CreateUserRequest
		locationHeader      string
		saveUserCallsAmount int
		expectedStatus      int
	}{
		{
			name: "success",
			request: CreateUserRequest{
				Name:     "Test",
				Username: "test",
				Email:    "test@mail.com",
			},
			locationHeader:      "/users/1",
			saveUserCallsAmount: 1,
			expectedStatus:      http.StatusCreated,
		},
		{
			name: "empty name",
			request: CreateUserRequest{
				Name:     "",
				Username: "test",
				Email:    "test@mail.com",
			},
			locationHeader:      "",
			saveUserCallsAmount: 0,
			expectedStatus:      http.StatusBadRequest,
		},
		{
			name: "empty username",
			request: CreateUserRequest{
				Name:     "Test",
				Username: "",
				Email:    "test@mail.com",
			},
			locationHeader:      "",
			saveUserCallsAmount: 0,
			expectedStatus:      http.StatusBadRequest,
		},
		{
			name: "invalid email",
			request: CreateUserRequest{
				Name:     "Test",
				Username: "test",
				Email:    "test",
			},
			locationHeader:      "",
			saveUserCallsAmount: 0,
			expectedStatus:      http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &userservice.MockUserRepository{
				SaveFunc: func(_ *userentity.User) error {
					return nil
				},
			}
			bitcoinRepository := &bitcoinservice.MockBTCRepository{}

			service := userservice.NewUserService(userRepository, bitcoinRepository)

			sut := NewUserHTTPHandlers(service, getUserIDFromURL).CreateUser

			tests.HTTPExpect(t, sut).
				POST("/users/").
				WithJSON(tc.request).
				Expect().
				Status(tc.expectedStatus).
				ContentType("application/json", "utf-8").
				Header("Location")

			assert.Len(t, userRepository.SaveCalls(), tc.saveUserCallsAmount)
		})
	}
}
