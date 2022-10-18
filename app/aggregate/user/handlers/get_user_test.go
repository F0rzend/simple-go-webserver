package userhandlers

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	userentity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"

	bitcoinservice "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	userservice "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"

	"github.com/F0rzend/simple-go-webserver/app/tests"
)

func TestUserHTTPHandlers_GetUser(t *testing.T) {
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
				0,
				time.Now(),
				time.Now(),
			)
		},
	}
	bitcoinRepository := &bitcoinservice.MockBTCRepository{}

	service := userservice.NewUserService(userRepository, bitcoinRepository)

	sut := NewUserHTTPHandlers(service, getUserIDFromURL).GetUser

	tests.HTTPExpect(t, sut).
		GET("/users/1").
		Expect().
		ContentType("application/json", "utf-8").
		Status(expectedStatus)

	assert.Len(t, userRepository.GetCalls(), 1)
}
