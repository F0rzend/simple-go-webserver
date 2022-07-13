package handlers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	userService "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
	"github.com/F0rzend/simple-go-webserver/app/tests"
)

func TestUserHTTPHandlers_ChangeBTCBalance(t *testing.T) {
	t.Parallel()

	getUserIDFromURL := func(r *http.Request) (uint64, error) {
		return 1, nil
	}

	service := &userService.MockUserService{
		ChangeBitcoinBalanceFunc: func(command userService.ChangeBitcoinBalanceCommand) error {
			return nil
		},
	}

	handler := http.HandlerFunc(NewUserHTTPHandlers(service, getUserIDFromURL).ChangeBTCBalance)

	w, r := tests.PrepareHandlerArgs(t, http.MethodPost, "/users/1/bitcoin", ChangeBTCBalanceRequest{
		Action: "buy",
		Amount: 1,
	})
	handler.ServeHTTP(w, r)

	tests.AssertStatus(t, w, r, http.StatusNoContent)
	assert.Len(t, service.ChangeBitcoinBalanceCalls(), 1)
}
