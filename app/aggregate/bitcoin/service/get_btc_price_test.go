package bitcoinservice

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
)

func TestBitcoinService_GetBTCPrice(t *testing.T) {
	t.Parallel()

	priceInUSD, _ := bitcoinentity.NewUSD(1)
	bitcoinPrice := bitcoinentity.NewBTCPrice(
		priceInUSD,
		time.Now(),
	)
	bitcoinRepository := &bitcoinrepositories.MockBTCRepository{
		GetPriceFunc: func() bitcoinentity.BTCPrice {
			return bitcoinPrice
		},
	}

	service := NewBitcoinService(bitcoinRepository)

	actual := service.GetBTCPrice()

	assert.Equal(t, bitcoinPrice, actual)
	assert.Len(t, bitcoinRepository.GetPriceCalls(), 1)
}
