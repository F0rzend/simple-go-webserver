package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	bitcoinRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
)

func TestBitcoinService_SetBTCPrice(t *testing.T) {
	t.Parallel()

	bitcoinRepository := &bitcoinRepositories.MockBTCRepository{
		SetPriceFunc: func(price bitcoinEntity.USD) error {
			return nil
		},
	}

	service := NewBitcoinService(bitcoinRepository)

	err := service.SetBTCPrice(1)

	assert.NoError(t, err)
}
