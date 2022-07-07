package handlers

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	"github.com/F0rzend/simple-go-webserver/app/tests"
)

// Tests BitcoinHTTPHandlers.getBTCPrice (GET /bitcoin endpoint)
func TestServer_GetBTCPrice(t *testing.T) {
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

// Tests BitcoinHTTPHandlers.setBTCPrice (POST /bitcoin endpoint)
func TestServer_SetBTCPrice(t *testing.T) {
	t.Parallel()

	server := getHTTPHandler(t)

	testCases := []struct {
		name     string
		body     SetBTCPriceRequest
		expected int
	}{
		{
			name: "success",
			body: SetBTCPriceRequest{
				Price: 1,
			},
			expected: http.StatusNoContent,
		},
		{
			name: "invalid amount",
			body: SetBTCPriceRequest{
				Price: -1,
			},
			expected: http.StatusBadRequest,
		},
		{
			name:     "empty body",
			body:     SetBTCPriceRequest{},
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

	bitcoinService := service.NewBitcoinService(
		tests.NewMockBitcoinRepository(),
	)
	handlers := NewBitcoinHTTPHandlers(bitcoinService)

	r := chi.NewRouter()
	handlers.SetRoutes(r)

	return r
}
