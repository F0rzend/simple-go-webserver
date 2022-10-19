package userhandlers

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"

	userentity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"

	bitcoinservice "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	userservice "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"

	"github.com/F0rzend/simple-go-webserver/app/tests"
)

func TestUserHTTPHandlers_GetUserBalance(t *testing.T) {
	t.Parallel()

	const expectedStatus = http.StatusOK

	getUserIDFromURL := func(_ *http.Request) (uint64, error) {
		return 1, nil
	}

	userRepository := &userservice.MockUserRepository{
		GetFunc: func(id uint64) (*userentity.User, error) {
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
		},
	}
	bitcoinRepository := &bitcoinservice.MockBTCRepository{
		GetPriceFunc: func() bitcoinentity.BTCPrice {
			return bitcoinentity.NewBTCPrice(bitcoinentity.NewUSD(1), time.Now())
		},
	}

	service := userservice.NewUserService(userRepository, bitcoinRepository)

	sut := NewUserHTTPHandlers(service, getUserIDFromURL).GetUserBalance

	tests.HTTPExpect(t, sut).
		POST("/users/1/balance").
		Expect().
		Status(expectedStatus).
		ContentType("application/json", "utf-8").
		JSON().Object().ValueEqual("balance", "100")

	assert.Len(t, userRepository.GetCalls(), 1)
}
