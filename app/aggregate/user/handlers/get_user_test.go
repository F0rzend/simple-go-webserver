package userhandlers

import (
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	bitcoinservice "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	userentity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userservice "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
	"github.com/F0rzend/simple-go-webserver/app/tests"
)

func TestUserHTTPHandlers_GetUser(t *testing.T) {
	t.Parallel()

	now := time.Now()
	const expectedStatus = http.StatusOK

	getUserFunc := func(id uint64) (*userentity.User, error) {
		return userentity.NewUser(
			1,
			"John",
			"john",
			"john@mail.com",
			0,
			0,
			now,
			now,
		)
	}
	getUserIDFromURL := func(_ *http.Request) (uint64, error) {
		return 1, nil
	}

	userRepository := &userservice.MockUserRepository{
		GetFunc: getUserFunc,
	}
	bitcoinRepository := &bitcoinservice.MockBTCRepository{}
	service := userservice.NewUserService(userRepository, bitcoinRepository)
	sut := NewUserHTTPHandlers(service, getUserIDFromURL).GetUser

	tests.HTTPExpect(t, sut).
		GET("/").
		Expect().
		Status(expectedStatus).
		ContentType("application/json", "utf-8").
		JSON().Object().Equal(
		UserResponse{
			ID:         1,
			Name:       "John",
			Username:   "john",
			Email:      "john@mail.com",
			BTCBalance: big.NewFloat(0),
			USDBalance: big.NewFloat(0),
			CreatedAt:  now,
			UpdatedAt:  now,
		})

	assert.Len(t, userRepository.GetCalls(), 1)
}
