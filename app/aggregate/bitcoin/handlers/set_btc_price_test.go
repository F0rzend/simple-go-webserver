package bitcoinhandlers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	bitcoinentity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"

	bitcoinservice "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"

	"github.com/F0rzend/simple-go-webserver/app/tests"
)

func TestSetBTCPrice(t *testing.T) {
	t.Parallel()

	request := SetBTCPriceRequest{Price: 1}
	const expectedStatus = http.StatusNoContent

	repository := &bitcoinservice.MockBTCRepository{
		SetPriceFunc: func(_ bitcoinentity.USD) error { return nil },
	}
	service := bitcoinservice.NewBitcoinService(repository)

	sut := NewBitcoinHTTPHandlers(service).SetBTCPrice

	tests.HTTPExpect(t, sut).
		POST("/bitcoin").
		WithJSON(request).
		Expect().
		Status(expectedStatus).
		ContentType("application/json", "utf-8")

	assert.Len(t, repository.SetPriceCalls(), 1)
}
