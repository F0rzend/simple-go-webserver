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

func TestUserHTTPHandlers_ChangeUSDBalance(t *testing.T) {
	t.Parallel()

	request := ChangeUSDBalanceRequest{
		Action: "withdraw",
		Amount: 1,
	}
	expectedStatus := http.StatusNoContent

	userRepository := &userservice.MockUserRepository{
		GetFunc: func(id uint64) (*userentity.User, error) {
			return userentity.NewUser(
				id,
				"John",
				"john",
				"john@mail.com",
				0,
				1,
				time.Now(),
				time.Now(),
			)
		},
		SaveFunc: func(_ *userentity.User) error {
			return nil
		},
	}
	bitcoinRepository := &bitcoinservice.MockBTCRepository{}

	service := userservice.NewUserService(userRepository, bitcoinRepository)

	sut := NewUserHTTPHandlers(service, func(_ *http.Request) (uint64, error) {
		return 1, nil
	}).ChangeUSDBalance

	tests.HTTPExpect(t, sut).
		POST("/").
		WithJSON(request).
		Expect().
		Status(expectedStatus).
		ContentType("application/json", "utf-8").
		Header("Location").Equal("/users/1")

	assert.Len(t, userRepository.GetCalls(), 1)
	assert.Len(t, userRepository.SaveCalls(), 1)
}
