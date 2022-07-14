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

	getUserIDFromURL := func(_ *http.Request) (uint64, error) {
		return 1, nil
	}
	changeBitcoinBalanceFunc := func(_ userService.ChangeBitcoinBalanceCommand) error {
		return nil
	}

	testCases := []struct {
		name                            string
		request                         ChangeBTCBalanceRequest
		shouldContainLocationHeader     bool
		changeBitcoinBalanceCallsAmount int
		expectedStatus                  int
	}{
		{
			name: "success",
			request: ChangeBTCBalanceRequest{
				Action: "buy",
				Amount: 1,
			},
			shouldContainLocationHeader:     true,
			changeBitcoinBalanceCallsAmount: 1,
			expectedStatus:                  http.StatusNoContent,
		},
		{
			name: "invalid action",
			request: ChangeBTCBalanceRequest{
				Action: "invalid",
				Amount: 1,
			},
			shouldContainLocationHeader:     false,
			changeBitcoinBalanceCallsAmount: 0,
			expectedStatus:                  http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			service := &userService.MockUserService{
				ChangeBitcoinBalanceFunc: changeBitcoinBalanceFunc,
			}

			handler := http.HandlerFunc(NewUserHTTPHandlers(service, getUserIDFromURL).ChangeBTCBalance)

			w, r := tests.PrepareHandlerArgs(t, http.MethodPost, "/users/1/bitcoin", tc.request)
			handler.ServeHTTP(w, r)

			tests.AssertStatus(t, w, r, tc.expectedStatus)
			if tc.shouldContainLocationHeader {
				assert.Equal(t, "/users/1", w.Header().Get("Location"))
			}
			assert.Len(t, service.ChangeBitcoinBalanceCalls(), tc.changeBitcoinBalanceCallsAmount)
		})
	}
}
