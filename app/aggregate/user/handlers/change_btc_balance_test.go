package userhandlers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/F0rzend/simple-go-webserver/app/tests"
)

func TestUserHTTPHandlers_ChangeBTCBalance(t *testing.T) {
	t.Parallel()

	request := ChangeBTCBalanceRequest{
		Action: "buy",
		Amount: 1,
	}
	expectedStatus := http.StatusNoContent

	service := &MockUserService{
		ChangeBitcoinBalanceFunc: func(_ uint64, _ string, _ float64) error {
			return nil
		},
	}

	handler := http.HandlerFunc(NewUserHTTPHandlers(service, func(_ *http.Request) (uint64, error) {
		return 1, nil
	}).ChangeBTCBalance)

	w, r := tests.PrepareHandlerArgs(t, http.MethodPost, "/users/1/bitcoin", request)
	handler.ServeHTTP(w, r)

	tests.AssertStatus(t, w, r, expectedStatus)
	assert.Equal(t, "/users/1", w.Header().Get("Location"))
	assert.Len(t, service.ChangeBitcoinBalanceCalls(), 1)
}
