package endpoints

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"

	"github.com/F0rzend/simple-go-webserver/app/bitcoin/http/types"
	"github.com/F0rzend/simple-go-webserver/app/bitcoin/service"
	"github.com/F0rzend/simple-go-webserver/app/tests"
)

// Tests BitcoinHTTPEndpoints.GetBTC (GET /bitcoin endpoint)
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

// Tests BitcoinHTTPEndpoints.SetBTCPrice (POST /bitcoin endpoint)
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

	bitcoinService, err := service.NewComponentTestBitcoinService()
	if err != nil {
		t.Fatal(err)
	}

	r := chi.NewRouter()
	NewBitcoinHTTPEndpoints(bitcoinService).Register(r)

	return r
}
