package bitcoinhandlers

import (
	"net/http"
	"testing"
	"time"

	"github.com/F0rzend/simple-go-webserver/app/common"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	bitcoinservice "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	"github.com/F0rzend/simple-go-webserver/app/tests"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func TestGetBTCPrice(t *testing.T) {
	t.Parallel()

	const expectedStatus = http.StatusOK

	now := time.Now()

	repository := &bitcoinservice.MockBTCRepository{
		GetPriceFunc: func() bitcoinentity.BTCPrice {
			price, err := bitcoinentity.NewBTCPrice(bitcoinentity.NewUSD(1), now)
			require.NoError(t, err)
			return price
		},
	}
	service := bitcoinservice.NewBitcoinService(repository)
	handler := NewBitcoinHTTPHandlers(service).GetBTCPrice
	sut := common.ErrorHandler(handler)

	tests.HTTPExpect(t, sut).
		GET("/bitcoin").
		Expect().
		Status(expectedStatus).
		ContentType("application/json").
		JSON().Object().Value("btc").Object().
		ValueEqual("price", "1").
		ValueEqual("updated_at", now)

	assert.Len(t, repository.GetPriceCalls(), 1)
}
