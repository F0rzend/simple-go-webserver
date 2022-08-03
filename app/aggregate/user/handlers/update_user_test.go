package handlers

import (
	"net/http"
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/tests"
	"github.com/stretchr/testify/assert"
)

func TestUserHTTPHandlers_UpdateUser(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                        string
		request                     UpdateUserRequest
		shouldContainLocationHeader bool
		updateUserCallsAmount       int
		expectedStatus              int
	}{
		{
			name:                        "empty request",
			request:                     UpdateUserRequest{},
			shouldContainLocationHeader: false,
			updateUserCallsAmount:       0,
			expectedStatus:              http.StatusBadRequest,
		},
		{
			name: "update name and email",
			request: UpdateUserRequest{
				Name:  strPointer("test"),
				Email: strPointer("test@mail.com"),
			},
			shouldContainLocationHeader: true,
			updateUserCallsAmount:       1,
			expectedStatus:              http.StatusNoContent,
		},
		{
			name: "correct email",
			request: UpdateUserRequest{
				Email: strPointer("test@m"),
			},
			shouldContainLocationHeader: true,
			updateUserCallsAmount:       1,
			expectedStatus:              http.StatusNoContent,
		},
		{
			name: "incorrect email",
			request: UpdateUserRequest{
				Email: strPointer("test"),
			},
			shouldContainLocationHeader: false,
			updateUserCallsAmount:       0,
			expectedStatus:              http.StatusBadRequest,
		},
		{
			name: "update name",
			request: UpdateUserRequest{
				Name: strPointer("test"),
			},
			shouldContainLocationHeader: true,
			updateUserCallsAmount:       1,
			expectedStatus:              http.StatusNoContent,
		},
	}

	getUserIDFromURL := func(_ *http.Request) (uint64, error) {
		return 1, nil
	}
	updateUserFunc := func(_ uint64, _, _ *string) error {
		return nil
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			service := &MockUserService{UpdateUserFunc: updateUserFunc}

			handler := http.HandlerFunc(NewUserHTTPHandlers(service, getUserIDFromURL).UpdateUser)

			w, r := tests.PrepareHandlerArgs(t, http.MethodPost, "/users/1", tc.request)
			handler.ServeHTTP(w, r)

			tests.AssertStatus(t, w, r, tc.expectedStatus)
			if tc.shouldContainLocationHeader {
				assert.Equal(t, "/users/1", w.Header().Get("Location"))
			}
			assert.Len(t, service.UpdateUserCalls(), tc.updateUserCallsAmount)
		})
	}
}

func strPointer(s string) *string {
	return &s
}
