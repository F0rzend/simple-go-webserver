package handlers

import (
	"net/http"
	"testing"

	userService "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
	"github.com/F0rzend/simple-go-webserver/app/tests"
	"github.com/stretchr/testify/assert"
)

func TestUserHTTPHandlers_CreateUser(t *testing.T) {
	t.Parallel()

	getUserIDFromURL := func(_ *http.Request) (uint64, error) {
		return 1, nil
	}
	createUserFunc := func(_ userService.CreateUserCommand) (uint64, error) {
		return 1, nil
	}

	testCases := []struct {
		name                        string
		request                     CreateUserRequest
		shouldContainLocationHeader bool
		createUserCallsAmount       int
		expectedStatus              int
	}{
		{
			name: "success",
			request: CreateUserRequest{
				Name:     "Test",
				Username: "test",
				Email:    "test@mail.com",
			},
			shouldContainLocationHeader: true,
			createUserCallsAmount:       1,
			expectedStatus:              http.StatusCreated,
		},
		{
			name: "empty name",
			request: CreateUserRequest{
				Name:     "",
				Username: "test",
				Email:    "test@mail.com",
			},
			shouldContainLocationHeader: false,
			createUserCallsAmount:       0,
			expectedStatus:              http.StatusBadRequest,
		},
		{
			name: "empty username",
			request: CreateUserRequest{
				Name:     "Test",
				Username: "",
				Email:    "test@mail.com",
			},
			shouldContainLocationHeader: false,
			createUserCallsAmount:       0,
			expectedStatus:              http.StatusBadRequest,
		},
		{
			name: "invalid email",
			request: CreateUserRequest{
				Name:     "Test",
				Username: "test",
				Email:    "test",
			},
			shouldContainLocationHeader: false,
			createUserCallsAmount:       0,
			expectedStatus:              http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			service := &userService.MockUserService{CreateUserFunc: createUserFunc}

			handler := http.HandlerFunc(NewUserHTTPHandlers(service, getUserIDFromURL).CreateUser)

			w, r := tests.PrepareHandlerArgs(t, http.MethodPost, "/users/", tc.request)
			handler.ServeHTTP(w, r)

			tests.AssertStatus(t, w, r, tc.expectedStatus)
			if tc.shouldContainLocationHeader {
				assert.Equal(t, "/users/1", w.Header().Get("Location"))
			}
			assert.Len(t, service.CreateUserCalls(), tc.createUserCallsAmount)
		})
	}
}
