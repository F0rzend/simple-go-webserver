package handlers

import (
	"net/http"
	"testing"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"

	"github.com/F0rzend/simple-go-webserver/app/tests"
	"github.com/stretchr/testify/assert"
)

func TestUserHTTPHandlers_GetUserBalance(t *testing.T) {
	t.Parallel()

	getUserIDFromURL := func(_ *http.Request) (uint64, error) {
		return 1, nil
	}

	service := &MockUserService{
		GetUserBalanceFunc: func(_ uint64) (bitcoinEntity.USD, error) {
			return bitcoinEntity.USD{}, nil
		},
	}

	handler := http.HandlerFunc(NewUserHTTPHandlers(service, getUserIDFromURL).GetUserBalance)

	w, r := tests.PrepareHandlerArgs(t, http.MethodPost, "/users/1/balance", nil)
	handler.ServeHTTP(w, r)

	tests.AssertStatus(t, w, r, http.StatusOK)
	assert.Len(t, service.GetUserBalanceCalls(), 1)
}
