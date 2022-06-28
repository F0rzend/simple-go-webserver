package handlers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/F0rzend/simple-go-webserver/app/tests"
)

// Tests UserHTTPHandlers.createUser (POST /users endpoint)
func TestServer_CreateUser(t *testing.T) {
	t.Parallel()

	server := getHTTPHandler(t)

	testCases := []struct {
		name               string
		body               CreateUserRequest
		hasLocationHeader  bool
		expectedStatusCode int
	}{
		{
			name: "success",
			body: CreateUserRequest{
				Name:     "Test User",
				Username: "test",
				Email:    "testuser@example.com",
			},
			hasLocationHeader:  true,
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name: "invalid email",
			body: CreateUserRequest{
				Name:     "Test User",
				Username: "test",
				Email:    "",
			},
			hasLocationHeader:  false,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			w, _ := tests.AssertStatus(t, server, http.MethodPost, "/users", tc.body, tc.expectedStatusCode)
			if tc.hasLocationHeader {
				assert.Equal(t, "/users/1", w.Header().Get("Location"))
			}
		})
	}
}

// Tests UserHTTPHandlers.getUser (GET /users/{id} endpoint)
func TestServer_GetUser(t *testing.T) {
	t.Parallel()

	server := getHTTPHandler(t)

	testCases := []struct {
		name     string
		userID   int
		expected int
	}{
		{
			name:     "success",
			userID:   1,
			expected: http.StatusOK,
		},
		{
			name:     "not found",
			userID:   0,
			expected: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tests.AssertStatus(t, server, http.MethodGet, fmt.Sprintf("/users/%d", tc.userID), nil, tc.expected)
		})
	}
}

// Tests UserHTTPHandlers.updateUser (PUT /users/{id} endpoint)
func TestServer_UpdateUser(t *testing.T) {
	t.Parallel()

	server := getHTTPHandler(t)

	testCases := []struct {
		name               string
		body               UpdateUserRequest
		hasLocationHeader  bool
		expectedStatusCode int
	}{
		{
			name: "success",
			body: UpdateUserRequest{
				Name:  strPointer("Test User"),
				Email: strPointer("testuser@example.com"),
			},
			hasLocationHeader:  true,
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name: "invalid email",
			body: UpdateUserRequest{
				Name:  strPointer("Test User"),
				Email: strPointer("invalid"),
			},
			hasLocationHeader:  false,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "empty body",
			body: UpdateUserRequest{
				Name:  nil,
				Email: nil,
			},
			hasLocationHeader:  false,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			w, _ := tests.AssertStatus(t, server, http.MethodPut, "/users/1", tc.body, tc.expectedStatusCode)
			if tc.hasLocationHeader {
				assert.Equal(t, "/users/1", w.Header().Get("Location"))
			}
		})
	}
}

// Tests UserHTTPHandlers.getUserBalance (Get /users/{userID}/balance endpoint)
func TestServer_GetUserBalance(t *testing.T) {
	t.Parallel()

	server := getHTTPHandler(t)

	testCases := []struct {
		name     string
		userID   int
		expected int
	}{
		{
			name:     "success",
			userID:   1,
			expected: http.StatusOK,
		},
		{
			name:     "not found",
			userID:   0,
			expected: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tests.AssertStatus(t, server, http.MethodGet, fmt.Sprintf("/users/%d/balance", tc.userID), nil, tc.expected)
		})
	}
}

// Tests UserHTTPHandlers.changeUSDBalance (POST /users/{userID}/usd endpoint)
func TestServer_ChangeUSDBalance(t *testing.T) {
	t.Parallel()

	server := getHTTPHandler(t)

	testCases := []struct {
		name     string
		body     ChangeUSDBalanceRequest
		userID   int
		expected int
	}{
		{
			name: "success",
			body: ChangeUSDBalanceRequest{
				Action: "deposit",
				Amount: 1,
			},
			userID:   1,
			expected: http.StatusNoContent,
		},
		{
			name: "invalid amount",
			body: ChangeUSDBalanceRequest{
				Action: "deposit",
				Amount: -1,
			},
			userID:   1,
			expected: http.StatusBadRequest,
		},
		{
			name: "insufficient funds",
			body: ChangeUSDBalanceRequest{
				Action: "withdraw",
				Amount: 1,
			},
			userID:   1,
			expected: http.StatusBadRequest,
		},
		{
			name:     "empty body",
			body:     ChangeUSDBalanceRequest{},
			userID:   1,
			expected: http.StatusBadRequest,
		},
		{
			name: "not found",
			body: ChangeUSDBalanceRequest{
				Action: "deposit",
				Amount: 1,
			},
			userID:   0,
			expected: http.StatusNotFound,
		},
		{
			name: "invalid action",
			body: ChangeUSDBalanceRequest{
				Action: "invalid",
				Amount: 1,
			},
			userID:   1,
			expected: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tests.AssertStatus(t, server, http.MethodPost, fmt.Sprintf("/users/%d/usd", tc.userID), tc.body, tc.expected)
		})
	}
}

// Tests UserHTTPHandlers.changeBTCBalance (POST /users/{userID}/bitcoin endpoint)
func TestServer_ChangeBTCBalance(t *testing.T) {
	t.Parallel()

	server := getHTTPHandler(t)

	testCases := []struct {
		name     string
		body     ChangeBTCBalanceRequest
		userID   int
		expected int
	}{
		{
			name: "success",
			body: ChangeBTCBalanceRequest{
				Action: "buy",
				Amount: 1,
			},
			userID:   2,
			expected: http.StatusNoContent,
		},
		{
			name: "invalid amount",
			body: ChangeBTCBalanceRequest{
				Action: "buy",
				Amount: -1,
			},
			userID:   1,
			expected: http.StatusBadRequest,
		},
		{
			name: "insufficient funds",
			body: ChangeBTCBalanceRequest{
				Action: "sell",
				Amount: 1,
			},
			userID:   1,
			expected: http.StatusBadRequest,
		},
		{
			name:     "empty body",
			body:     ChangeBTCBalanceRequest{},
			userID:   1,
			expected: http.StatusBadRequest,
		},
		{
			name: "not found",
			body: ChangeBTCBalanceRequest{
				Action: "buy",
				Amount: 1,
			},
			userID:   0,
			expected: http.StatusNotFound,
		},
		{
			name: "invalid action",
			body: ChangeBTCBalanceRequest{
				Action: "invalid",
				Amount: 1,
			},
			userID:   1,
			expected: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tests.AssertStatus(t, server, http.MethodPost, fmt.Sprintf("/users/%d/bitcoin", tc.userID), tc.body, tc.expected)
		})
	}
}

func getHTTPHandler(t *testing.T) http.Handler {
	t.Helper()

	userService := service.NewUserService(
		tests.NewMockUserRepository(),
		tests.NewMockBitcoinRepository(),
	)

	r := chi.NewRouter()
	NewUserHTTPHandlers(userService).SetRoutes(r)

	return r
}

func strPointer(s string) *string {
	return &s
}
