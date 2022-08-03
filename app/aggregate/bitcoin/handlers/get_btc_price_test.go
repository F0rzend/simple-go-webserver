package bitcoinhandlers

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/F0rzend/simple-go-webserver/app/tests"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func TestGetBTCPrice(t *testing.T) {
	t.Parallel()

	service := &MockBitcoinService{
		GetBTCPriceFunc: func() bitcoinentity.BTCPrice {
			btcPrice, _ := bitcoinentity.NewUSD(1)
			return bitcoinentity.NewBTCPrice(
				btcPrice,
				time.Now(),
			)
		},
	}
	handler := http.HandlerFunc(
		NewBitcoinHTTPHandlers(service).GetBTCPrice,
	)

	w, r := tests.PrepareHandlerArgs(t, http.MethodGet, "/bitcoin", nil)
	handler.ServeHTTP(w, r)
	tests.AssertStatus(t, w, r, http.StatusOK)
	assert.Len(t, service.GetBTCPriceCalls(), 1)
}
