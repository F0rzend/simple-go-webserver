package bitcoinservice

import (
	"net/http"
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/common"

	"github.com/stretchr/testify/assert"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
)

func TestBitcoinService_SetBTCPrice(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		price float64
		err   error
	}{
		{
			name:  "success",
			price: 1.0,
			err:   nil,
		},
		{
			name:  "negative price",
			price: -1.0,
			err: common.NewApplicationError(
				http.StatusBadRequest,
				"The amount of currency cannot be negative. Please pass a number greater than 0",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			bitcoinRepository := &bitcoinrepositories.MockBTCRepository{
				SetPriceFunc: func(price bitcoinentity.USD) error {
					return nil
				},
			}

			service := NewBitcoinService(bitcoinRepository)

			err := service.SetBTCPrice(tc.price)

			assert.Equal(t, tc.err, err)
		})
	}
}
