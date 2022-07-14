package handlers

import (
	"net/http"
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/tests"
	"github.com/stretchr/testify/assert"

	bitcoinService "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
)

func TestSetBTCPrice(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name               string
		newPrice           float64
		hasLocationHeader  bool
		serviceCallsAmount int
		expectedStatusCode int
	}{
		{
			name:               "success",
			newPrice:           1.0,
			hasLocationHeader:  true,
			serviceCallsAmount: 1,
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name:               "negative",
			newPrice:           -1.0,
			hasLocationHeader:  false,
			serviceCallsAmount: 0,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			service := &bitcoinService.MockBitcoinService{
				SetBTCPriceFunc: func(newPrice float64) error {
					return nil
				},
			}

			handler := http.HandlerFunc(
				NewBitcoinHTTPHandlers(service).SetBTCPrice,
			)

			w, r := tests.PrepareHandlerArgs(t,
				http.MethodPost,
				"/bitcoin",
				SetBTCPriceRequest{Price: tc.newPrice},
			)
			handler.ServeHTTP(w, r)

			tests.AssertStatus(t, w, r, tc.expectedStatusCode)
			if tc.hasLocationHeader {
				assert.Equal(t, w.Header().Get("Location"), "/bitcoin")
			}
			assert.Len(t, service.SetBTCPriceCalls(), tc.serviceCallsAmount)
		})
	}
}
