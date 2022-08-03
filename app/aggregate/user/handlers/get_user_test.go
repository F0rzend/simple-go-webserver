package handlers

import (
	"net/http"
	"net/mail"
	"testing"

	userEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"

	"github.com/F0rzend/simple-go-webserver/app/tests"
	"github.com/stretchr/testify/assert"
)

func TestUserHTTPHandlers_GetUser(t *testing.T) {
	t.Parallel()

	getUserIDFromURL := func(_ *http.Request) (uint64, error) {
		return 1, nil
	}

	service := &MockUserService{
		GetUserFunc: func(_ uint64) (*userEntity.User, error) {
			return &userEntity.User{
				Email: &mail.Address{Address: "test@mail.com"},
			}, nil
		},
	}

	handler := http.HandlerFunc(NewUserHTTPHandlers(service, getUserIDFromURL).GetUser)

	w, r := tests.PrepareHandlerArgs(t, http.MethodGet, "/users/1", nil)
	handler.ServeHTTP(w, r)

	tests.AssertStatus(t, w, r, http.StatusOK)
	assert.Len(t, service.GetUserCalls(), 1)
}
