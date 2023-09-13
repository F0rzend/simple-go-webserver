package userhandlers

import (
	"net/http"
	"testing"
	"time"

	"github.com/F0rzend/simple-go-webserver/app/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	bitcoinservice "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	userentity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userservice "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
	"github.com/F0rzend/simple-go-webserver/app/tests"
)

func TestUserHTTPHandlers_GetUserBalance(t *testing.T) {
	t.Parallel()

	const expectedStatus = http.StatusOK

	getUserFunc := func(id uint64) (*userentity.User, error) {
		return userentity.NewUser(
			id,
			"John",
			"john",
			"john@mail.com",
			0,
			100,
			time.Now(),
			time.Now(),
		)
	}
	getPriceFunc := func() bitcoinentity.BTCPrice {
		price, err := bitcoinentity.NewBTCPrice(bitcoinentity.NewUSD(1), time.Now())
		require.NoError(t, err)

		return price
	}
	getUserIDFromURL := func(_ *http.Request) (uint64, error) {
		return 1, nil
	}

	userRepository := &userservice.MockUserRepository{
		GetFunc: getUserFunc,
	}
	bitcoinRepository := &bitcoinservice.MockBTCRepository{
		GetPriceFunc: getPriceFunc,
	}
	service := userservice.NewUserService(userRepository, bitcoinRepository)
	handler := NewUserHTTPHandlers(service, getUserIDFromURL).GetUserBalance
	sut := common.ErrorHandler(handler)

	tests.HTTPExpect(t, sut).
		POST("/").
		Expect().
		Status(expectedStatus).
		ContentType("application/json").
		JSON().Object().ValueEqual("balance", "100")

	assert.Len(t, userRepository.GetCalls(), 1)
	assert.Len(t, bitcoinRepository.GetPriceCalls(), 1)
}
