package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	bitcoinRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
)

func TestBitcoinService_GetBTCPrice(t *testing.T) {
	t.Parallel()

	priceInUSD, _ := bitcoinEntity.NewUSD(1)
	bitcoinPrice := bitcoinEntity.NewBTCPrice(
		priceInUSD,
		time.Now(),
	)
	bitcoinRepository := &bitcoinRepositories.MockBTCRepository{
		GetPriceFunc: func() bitcoinEntity.BTCPrice {
			return bitcoinPrice
		},
	}

	service := NewBitcoinService(bitcoinRepository)

	actual := service.GetBTCPrice()

	assert.Equal(t, bitcoinPrice, actual)
	assert.Len(t, bitcoinRepository.GetPriceCalls(), 1)
}
