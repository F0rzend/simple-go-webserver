package handlers

import (
	"net/http"
	"testing"

	userService "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
	"github.com/F0rzend/simple-go-webserver/app/tests"
	"github.com/stretchr/testify/assert"
)

func TestUserHTTPHandlers_ChangeUSDBalance(t *testing.T) {
	t.Parallel()

	getUserIDFromURL := func(_ *http.Request) (uint64, error) {
		return 1, nil
	}
	changeUserBalanceFunc := func(_ userService.ChangeUserBalanceCommand) error {
		return nil
	}

	testCases := []struct {
		name                         string
		request                      ChangeUSDBalanceRequest
		shouldContainLocationHeader  bool
		changeUserBalanceCallsAmount int
		expectedStatus               int
	}{
		{
			name: "success",
			request: ChangeUSDBalanceRequest{
				Action: "withdraw",
				Amount: 1,
			},
			shouldContainLocationHeader:  true,
			changeUserBalanceCallsAmount: 1,
			expectedStatus:               http.StatusNoContent,
		},
		{
			name: "invalid action",
			request: ChangeUSDBalanceRequest{
				Action: "invalid",
				Amount: 1,
			},
			shouldContainLocationHeader:  false,
			changeUserBalanceCallsAmount: 0,
			expectedStatus:               http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			service := &userService.MockUserService{
				ChangeUserBalanceFunc: changeUserBalanceFunc,
			}

			handler := http.HandlerFunc(NewUserHTTPHandlers(service, getUserIDFromURL).ChangeUSDBalance)

			w, r := tests.PrepareHandlerArgs(t, http.MethodPost, "/users/1/usd", tc.request)
			handler.ServeHTTP(w, r)

			tests.AssertStatus(t, w, r, tc.expectedStatus)
			if tc.shouldContainLocationHeader {
				assert.Equal(t, "/users/1", w.Header().Get("Location"))
			}
			assert.Len(t, service.ChangeUserBalanceCalls(), tc.changeUserBalanceCallsAmount)
		})
	}
}
