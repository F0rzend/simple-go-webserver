package bitcoinhandlers

import (
	"net/http"
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/common"

	bitcoinentity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	bitcoinservice "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	"github.com/F0rzend/simple-go-webserver/app/tests"
	"github.com/stretchr/testify/assert"
)

func TestSetBTCPrice(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                          string
		price                         float64
		expectedStatus                int
		repositorySetPriceCallsAmount int
	}{
		{
			name:                          "success",
			price:                         100.0,
			expectedStatus:                http.StatusNoContent,
			repositorySetPriceCallsAmount: 1,
		},
		{
			name:                          "empty price",
			price:                         0.0,
			expectedStatus:                http.StatusBadRequest,
			repositorySetPriceCallsAmount: 0,
		},
		{
			name:                          "negative price",
			price:                         -100.0,
			expectedStatus:                http.StatusBadRequest,
			repositorySetPriceCallsAmount: 0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			repository := &bitcoinservice.MockBTCRepository{
				SetPriceFunc: func(_ bitcoinentity.BTCPrice) error { return nil },
			}
			service := bitcoinservice.NewBitcoinService(repository)
			handler := NewBitcoinHTTPHandlers(service).SetBTCPrice
			sut := common.ErrorHandler(handler)

			tests.HTTPExpect(t, sut).
				POST("/bitcoin").
				WithJSON(SetBTCPriceRequest{Price: tc.price}).
				Expect().
				Status(tc.expectedStatus).
				ContentType("application/json")

			assert.Len(t, repository.SetPriceCalls(), tc.repositorySetPriceCallsAmount)
		})
	}
}
