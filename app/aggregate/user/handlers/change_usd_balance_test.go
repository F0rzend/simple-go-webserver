package handlers

import (
	"net/http"
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/tests"
	"github.com/stretchr/testify/assert"
)

func TestUserHTTPHandlers_ChangeUSDBalance(t *testing.T) {
	t.Parallel()

	request := ChangeUSDBalanceRequest{
		Action: "withdraw",
		Amount: 1,
	}
	expectedStatus := http.StatusNoContent

	service := &MockUserService{
		ChangeUserBalanceFunc: func(_ uint64, _ string, _ float64) error {
			return nil
		},
	}

	handler := http.HandlerFunc(NewUserHTTPHandlers(service, func(_ *http.Request) (uint64, error) {
		return 1, nil
	}).ChangeUSDBalance)

	w, r := tests.PrepareHandlerArgs(t, http.MethodPost, "/users/1/usd", request)
	handler.ServeHTTP(w, r)

	tests.AssertStatus(t, w, r, expectedStatus)
	assert.Equal(t, "/users/1", w.Header().Get("Location"))
	assert.Len(t, service.ChangeUserBalanceCalls(), 1)
}
