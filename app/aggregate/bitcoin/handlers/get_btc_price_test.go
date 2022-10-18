package bitcoinhandlers

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	bitcoinservice "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	"github.com/F0rzend/simple-go-webserver/app/tests"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func TestGetBTCPrice(t *testing.T) {
	t.Parallel()

	const expectedStatus = http.StatusOK

	repository := &bitcoinservice.MockBTCRepository{
		GetPriceFunc: func() bitcoinentity.BTCPrice {
			return bitcoinentity.NewBTCPrice(bitcoinentity.NewUSD(1), time.Now())
		},
	}
	service := bitcoinservice.NewBitcoinService(repository)

	sut := NewBitcoinHTTPHandlers(service).GetBTCPrice

	tests.HTTPExpect(t, sut).
		GET("/bitcoin").
		Expect().
		Status(expectedStatus).
		ContentType("application/json", "utf-8")

	assert.Len(t, repository.GetPriceCalls(), 1)
}
