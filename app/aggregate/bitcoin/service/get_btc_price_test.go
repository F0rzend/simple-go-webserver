package bitcoinservice

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func TestBitcoinService_GetBTCPrice(t *testing.T) {
	t.Parallel()

	bitcoinPrice, _ := bitcoinentity.NewBTCPrice(
		bitcoinentity.NewUSD(1),
		time.Now(),
	)
	bitcoinRepository := &MockBTCRepository{
		GetPriceFunc: func() bitcoinentity.BTCPrice {
			return bitcoinPrice
		},
	}
	sut := NewBitcoinService(bitcoinRepository)

	actual := sut.GetBTCPrice()

	assert.Equal(t, bitcoinPrice, actual)
	assert.Len(t, bitcoinRepository.GetPriceCalls(), 1)
}
