package bitcoinservice

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func TestBitcoinService_SetBTCPrice(t *testing.T) {
	t.Parallel()

	bitcoinRepository := &MockBTCRepository{
		SetPriceFunc: func(_ bitcoinentity.BTCPrice) error { return nil },
	}

	sut := NewBitcoinService(bitcoinRepository)

	err := sut.SetBTCPrice(1.0)
	assert.NoError(t, err)
}
