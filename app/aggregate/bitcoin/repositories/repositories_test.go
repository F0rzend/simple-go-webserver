package bitcoinrepositories

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	bitcoinentity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func TestMemoryBTCRepositories(t *testing.T) {
	t.Parallel()

	sut := NewMemoryBTCRepository()

	defaultPrice := sut.GetPrice()
	assert.True(t, defaultPrice.GetPrice().Equal(bitcoinentity.NewUSD(100.0)))

	price, _ := bitcoinentity.NewBTCPrice(bitcoinentity.NewUSD(100.0), time.Now())
	err := sut.SetPrice(price)
	assert.NoError(t, err)

	actualPrice := sut.GetPrice()
	assert.True(t, actualPrice.GetPrice().Equal(bitcoinentity.NewUSD(100.0)))
}
