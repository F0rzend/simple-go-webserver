package bitcoinservice

import (
	"net/http"
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/tests"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func TestBitcoinService_SetBTCPrice(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                string
		price               float64
		expectedErrorStatus int
	}{
		{
			name:                "success",
			price:               1.0,
			expectedErrorStatus: 0,
		},
		{
			name:                "negative",
			price:               -1.0,
			expectedErrorStatus: http.StatusBadRequest,
		},
	}

	bitcoinRepository := &MockBTCRepository{
		SetPriceFunc: func(_ bitcoinentity.USD) error { return nil },
	}

	sut := NewBitcoinService(bitcoinRepository)

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := sut.SetBTCPrice(tc.price)

			tests.ExpectApplicationError(t, tc.expectedErrorStatus, err)
		})
	}
}
