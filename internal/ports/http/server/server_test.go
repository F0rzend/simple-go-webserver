package server

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/F0rzend/SimpleGoWebserver/internal/tests"

	"github.com/F0rzend/SimpleGoWebserver/internal/application"

	"github.com/F0rzend/SimpleGoWebserver/internal/ports/http/types"

	"github.com/stretchr/testify/assert"
)

// Tests Server.CreateUser (POST /users endpoint)
func TestServer_CreateUser(t *testing.T) {
	t.Parallel()

	server := getHTTPHandler(t)

	testCases := []struct {
		name               string
		body               types.CreateUserRequest
		hasLocationHeader  bool
		expectedStatusCode int
	}{
		{
			name: "success",
			body: types.CreateUserRequest{
				Name:     "Test User",
				Username: "test",
				Email:    "testuser@example.com",
			},
			hasLocationHeader:  true,
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name: "invalid email",
			body: types.CreateUserRequest{
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

// Tests Server.GetUser (GET /users/{id} endpoint)
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

// Tests Server.UpdateUser (PUT /users/{id} endpoint)
func TestServer_UpdateUser(t *testing.T) {
	t.Parallel()

	server := getHTTPHandler(t)

	testCases := []struct {
		name               string
		body               types.UpdateUserRequest
		hasLocationHeader  bool
		expectedStatusCode int
	}{
		{
			name: "success",
			body: types.UpdateUserRequest{
				Name:  strPointer("Test User"),
				Email: strPointer("testuser@example.com"),
			},
			hasLocationHeader:  true,
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name: "invalid email",
			body: types.UpdateUserRequest{
				Name:  strPointer("Test User"),
				Email: strPointer("invalid"),
			},
			hasLocationHeader:  false,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "empty body",
			body: types.UpdateUserRequest{
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

// Tests Server.GetUserBalance (Get /users/{userID}/balance endpoint)
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

// Tests Server.ChangeUserBalance (POST /users/{userID}/usd endpoint)
func TestServer_ChangeUSDBalance(t *testing.T) {
	t.Parallel()

	server := getHTTPHandler(t)

	testCases := []struct {
		name     string
		body     types.ChangeUSDBalanceRequest
		userID   int
		expected int
	}{
		{
			name: "success",
			body: types.ChangeUSDBalanceRequest{
				Action: "deposit",
				Amount: 1,
			},
			userID:   1,
			expected: http.StatusNoContent,
		},
		{
			name: "invalid amount",
			body: types.ChangeUSDBalanceRequest{
				Action: "deposit",
				Amount: -1,
			},
			userID:   1,
			expected: http.StatusBadRequest,
		},
		{
			name: "insufficient funds",
			body: types.ChangeUSDBalanceRequest{
				Action: "withdraw",
				Amount: 1,
			},
			userID:   1,
			expected: http.StatusBadRequest,
		},
		{
			name:     "empty body",
			body:     types.ChangeUSDBalanceRequest{},
			userID:   1,
			expected: http.StatusBadRequest,
		},
		{
			name: "not found",
			body: types.ChangeUSDBalanceRequest{
				Action: "deposit",
				Amount: 1,
			},
			userID:   0,
			expected: http.StatusNotFound,
		},
		{
			name: "invalid action",
			body: types.ChangeUSDBalanceRequest{
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

// Tests Server.ChangeBTCBalance (POST /users/{userID}/bitcoin endpoint)
func TestServer_ChangeBTCBalance(t *testing.T) {
	t.Parallel()

	server := getHTTPHandler(t)

	testCases := []struct {
		name     string
		body     types.ChangeBTCBalanceRequest
		userID   int
		expected int
	}{
		{
			name: "success",
			body: types.ChangeBTCBalanceRequest{
				Action: "buy",
				Amount: 1,
			},
			userID:   2,
			expected: http.StatusNoContent,
		},
		{
			name: "invalid amount",
			body: types.ChangeBTCBalanceRequest{
				Action: "buy",
				Amount: -1,
			},
			userID:   1,
			expected: http.StatusBadRequest,
		},
		{
			name: "insufficient funds",
			body: types.ChangeBTCBalanceRequest{
				Action: "sell",
				Amount: 1,
			},
			userID:   1,
			expected: http.StatusBadRequest,
		},
		{
			name:     "empty body",
			body:     types.ChangeBTCBalanceRequest{},
			userID:   1,
			expected: http.StatusBadRequest,
		},
		{
			name: "not found",
			body: types.ChangeBTCBalanceRequest{
				Action: "buy",
				Amount: 1,
			},
			userID:   0,
			expected: http.StatusNotFound,
		},
		{
			name: "invalid action",
			body: types.ChangeBTCBalanceRequest{
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

// Tests Server.GetBTC (GET /bitcoin endpoint)
func TestServer_GetBTC(t *testing.T) {
	t.Parallel()

	server := getHTTPHandler(t)

	tests.AssertStatus(t,
		server,
		http.MethodGet,
		"/bitcoin",
		nil,
		http.StatusOK,
	)
}

// Tests Server.SetBTCPrice (POST /bitcoin endpoint)
func TestServer_SetBTCPrice(t *testing.T) {
	t.Parallel()

	server := getHTTPHandler(t)

	testCases := []struct {
		name     string
		body     types.SetBTCPriceRequest
		expected int
	}{
		{
			name: "success",
			body: types.SetBTCPriceRequest{
				Price: 1,
			},
			expected: http.StatusNoContent,
		},
		{
			name: "invalid amount",
			body: types.SetBTCPriceRequest{
				Price: -1,
			},
			expected: http.StatusBadRequest,
		},
		{
			name:     "empty body",
			body:     types.SetBTCPriceRequest{},
			expected: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tests.AssertStatus(t, server, http.MethodPut, "/bitcoin", tc.body, tc.expected)
		})
	}
}

func getHTTPHandler(t *testing.T) http.Handler {
	t.Helper()

	app, err := application.NewComponentTestApplication()
	if err != nil {
		t.Fatal(err)
	}
	return NewServer(app).GetRouter()
}

func strPointer(s string) *string {
	return &s
}
