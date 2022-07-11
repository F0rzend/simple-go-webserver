package handlers

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"

	"github.com/F0rzend/simple-go-webserver/app/tests"

	bitcoinService "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
)

func TestGetBTCPrice(t *testing.T) {
	t.Parallel()

	service := &bitcoinService.MockBitcoinService{
		GetBTCPriceFunc: func() bitcoinEntity.BTCPrice {
			btcPrice, _ := bitcoinEntity.NewUSD(1)
			return bitcoinEntity.NewBTCPrice(
				btcPrice,
				time.Now(),
			)
		},
	}
	handler := http.HandlerFunc(
		NewBitcoinHTTPHandlers(service).GetBTCPrice,
	)

	w, r := tests.PrepareHandlerArgs(t, http.MethodGet, "/bitcoin", nil)
	tests.ProcessHandler(t,
		handler,
		w, r,
		http.StatusOK,
	)
	assert.Len(t, service.GetBTCPriceCalls(), 1)
}
