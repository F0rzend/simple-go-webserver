package userhandlers

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	bitcoinservice "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	userentity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userservice "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
	"github.com/F0rzend/simple-go-webserver/app/tests"
)

func TestUserHTTPHandlers_ChangeBTCBalance(t *testing.T) {
	t.Parallel()

	type calls struct {
		getUser  int
		saveUser int
		getPrice int
	}

	type response struct {
		status   int
		location string
	}

	testCases := []struct {
		name     string
		request  ChangeBTCBalanceRequest
		response response
		calls    calls
	}{
		{
			name: "success buying",
			request: ChangeBTCBalanceRequest{
				Action: "buy",
				Amount: 10.0,
			},
			response: response{
				status:   http.StatusNoContent,
				location: "/users/1",
			},
			calls: calls{
				getUser:  1,
				saveUser: 1,
				getPrice: 1,
			},
		},
		{
			name: "success selling",
			request: ChangeBTCBalanceRequest{
				Action: "sell",
				Amount: 10.0,
			},
			response: response{
				status:   http.StatusNoContent,
				location: "/users/1",
			},
			calls: calls{
				getUser:  1,
				saveUser: 1,
				getPrice: 1,
			},
		},
		{
			name: "empty amount",
			request: ChangeBTCBalanceRequest{
				Action: "buy",
			},
			response: response{
				status: http.StatusBadRequest,
			},
		},
		{
			name: "empty action",
			request: ChangeBTCBalanceRequest{
				Amount: 10.0,
			},
			response: response{
				status: http.StatusBadRequest,
			},
		},
	}

	getUserFunc := func(id uint64) (*userentity.User, error) {
		return userentity.NewUser(
			id,
			"John",
			"john",
			"john@mail.com",
			10,
			10,
			time.Now(),
			time.Now(),
		)
	}
	getPriceFunc := func() bitcoinentity.BTCPrice {
		price, err := bitcoinentity.NewBTCPrice(bitcoinentity.NewUSD(1), time.Now())
		require.NoError(t, err)

		return price
	}
	saveUserFunc := func(_ *userentity.User) error {
		return nil
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository := &userservice.MockUserRepository{
				GetFunc:  getUserFunc,
				SaveFunc: saveUserFunc,
			}
			bitcoinRepository := &bitcoinservice.MockBTCRepository{GetPriceFunc: getPriceFunc}

			service := userservice.NewUserService(userRepository, bitcoinRepository)

			sut := NewUserHTTPHandlers(service, func(r *http.Request) (uint64, error) {
				return 1, nil
			}).ChangeBTCBalance

			tests.HTTPExpect(t, sut).
				POST("/").
				WithJSON(tc.request).
				Expect().
				Status(tc.response.status).
				ContentType("application/json", "utf-8").
				Header("Location").Equal(tc.response.location)

			assert.Len(t, userRepository.GetCalls(), tc.calls.getUser)
			assert.Len(t, bitcoinRepository.GetPriceCalls(), tc.calls.getPrice)
			assert.Len(t, userRepository.SaveCalls(), tc.calls.saveUser)
		})
	}
}
