package handlers

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stretchr/testify/assert"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	bitcoinRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	userRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
	userService "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
	"github.com/F0rzend/simple-go-webserver/app/tests"
)

// Tests UserHTTPHandlers.createUser (POST /users endpoint)
func TestServer_CreateUser(t *testing.T) {
	t.Parallel()

	handler := getHTTPHandler(t)

	testCases := []struct {
		name               string
		body               CreateUserRequest
		hasLocationHeader  bool
		expectedStatusCode int
	}{
		{
			name: "success",
			body: CreateUserRequest{
				Name:     "Test User",
				Username: "test",
				Email:    "testuser@example.com",
			},
			hasLocationHeader:  true,
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name: "invalid email",
			body: CreateUserRequest{
				Name:     "Test User",
				Username: "test",
				Email:    "",
			},
			hasLocationHeader:  false,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			w, r := tests.PrepareHandlerArgs(t, http.MethodPost, "/users", tc.body)
			tests.ProcessHandler(t, handler, w, r, tc.expectedStatusCode)
			if tc.hasLocationHeader {
				assert.Equal(t, "/users/1", w.Header().Get("Location"))
			}
		})
	}
}

// Tests UserHTTPHandlers.getUser (GET /users/{id} endpoint)
func TestServer_GetUser(t *testing.T) {
	t.Parallel()

	handler := getHTTPHandler(t)

	testCases := []struct {
		name               string
		userID             int
		expectedStatusCode int
	}{
		{
			name:               "success",
			userID:             1,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "not found",
			userID:             0,
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			url := fmt.Sprintf("/users/%d", tc.userID)
			w, r := tests.PrepareHandlerArgs(t, http.MethodGet, url, nil)
			tests.ProcessHandler(t, handler, w, r, tc.expectedStatusCode)
		})
	}
}

// Tests UserHTTPHandlers.updateUser (PUT /users/{id} endpoint)
func TestServer_UpdateUser(t *testing.T) {
	t.Parallel()

	handler := getHTTPHandler(t)

	testCases := []struct {
		name               string
		body               UpdateUserRequest
		hasLocationHeader  bool
		expectedStatusCode int
	}{
		{
			name: "success",
			body: UpdateUserRequest{
				Name:  strPointer("Test User"),
				Email: strPointer("testuser@example.com"),
			},
			hasLocationHeader:  true,
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name: "invalid email",
			body: UpdateUserRequest{
				Name:  strPointer("Test User"),
				Email: strPointer("invalid"),
			},
			hasLocationHeader:  false,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "empty body",
			body: UpdateUserRequest{
				Name:  nil,
				Email: nil,
			},
			hasLocationHeader:  false,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			w, r := tests.PrepareHandlerArgs(t, http.MethodPut, "/users/1", tc.body)
			tests.ProcessHandler(t, handler, w, r, tc.expectedStatusCode)
			if tc.hasLocationHeader {
				assert.Equal(t, "/users/1", w.Header().Get("Location"))
			}
		})
	}
}

// Tests UserHTTPHandlers.getUserBalance (Get /users/{userID}/balance endpoint)
func TestServer_GetUserBalance(t *testing.T) {
	t.Parallel()

	handler := getHTTPHandler(t)

	testCases := []struct {
		name               string
		userID             int
		expectedStatusCode int
	}{
		{
			name:               "success",
			userID:             1,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "not found",
			userID:             0,
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			url := fmt.Sprintf("/users/%d/balance", tc.userID)
			w, r := tests.PrepareHandlerArgs(t, http.MethodGet, url, nil)
			tests.ProcessHandler(t, handler, w, r, tc.expectedStatusCode)
		})
	}
}

// Tests UserHTTPHandlers.changeUSDBalance (POST /users/{userID}/usd endpoint)
func TestServer_ChangeUSDBalance(t *testing.T) {
	t.Parallel()

	handler := getHTTPHandler(t)

	testCases := []struct {
		name               string
		body               ChangeUSDBalanceRequest
		userID             int
		expectedStatusCode int
	}{
		{
			name: "success",
			body: ChangeUSDBalanceRequest{
				Action: "deposit",
				Amount: 1,
			},
			userID:             1,
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name: "invalid amount",
			body: ChangeUSDBalanceRequest{
				Action: "deposit",
				Amount: -1,
			},
			userID:             1,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "insufficient funds",
			body: ChangeUSDBalanceRequest{
				Action: "withdraw",
				Amount: 1,
			},
			userID:             1,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "empty body",
			body:               ChangeUSDBalanceRequest{},
			userID:             1,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "not found",
			body: ChangeUSDBalanceRequest{
				Action: "deposit",
				Amount: 1,
			},
			userID:             0,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: "invalid action",
			body: ChangeUSDBalanceRequest{
				Action: "invalid",
				Amount: 1,
			},
			userID:             1,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			url := fmt.Sprintf("/users/%d/usd", tc.userID)
			w, r := tests.PrepareHandlerArgs(t, http.MethodPost, url, tc.body)
			tests.ProcessHandler(t, handler, w, r, tc.expectedStatusCode)
		})
	}
}

// Tests UserHTTPHandlers.changeBTCBalance (POST /users/{userID}/bitcoin endpoint)
func TestServer_ChangeBTCBalance(t *testing.T) {
	t.Parallel()

	handler := getHTTPHandler(t)

	testCases := []struct {
		name               string
		body               ChangeBTCBalanceRequest
		userID             int
		expectedStatusCode int
	}{
		{
			name: "success",
			body: ChangeBTCBalanceRequest{
				Action: "buy",
				Amount: 1,
			},
			userID:             2,
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name: "invalid amount",
			body: ChangeBTCBalanceRequest{
				Action: "buy",
				Amount: -1,
			},
			userID:             1,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "insufficient funds",
			body: ChangeBTCBalanceRequest{
				Action: "sell",
				Amount: 1,
			},
			userID:             1,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "empty body",
			body:               ChangeBTCBalanceRequest{},
			userID:             1,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "not found",
			body: ChangeBTCBalanceRequest{
				Action: "buy",
				Amount: 1,
			},
			userID:             0,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: "invalid action",
			body: ChangeBTCBalanceRequest{
				Action: "invalid",
				Amount: 1,
			},
			userID:             1,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			url := fmt.Sprintf("/users/%d/bitcoin", tc.userID)
			w, r := tests.PrepareHandlerArgs(t, http.MethodPost, url, tc.body)
			tests.ProcessHandler(t, handler, w, r, tc.expectedStatusCode)
		})
	}
}

func getHTTPHandler(t *testing.T) http.Handler {
	t.Helper()

	handlers := NewUserHTTPHandlers(
		userService.NewUserService(
			NewMockUserRepository(),
			NewMockBitcoinRepository(),
		))

	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.AllowContentType("application/json"),
	)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", handlers.CreateUser)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetUser)
			r.Put("/", handlers.UpdateUser)
			r.Get("/balance", handlers.GetUserBalance)

			r.Post("/bitcoin", handlers.ChangeBTCBalance)
			r.Post("/usd", handlers.ChangeUSDBalance)
		})
	})

	return r
}

func strPointer(s string) *string {
	return &s
}

func NewMockUserRepository() entity.UserRepository {
	now := time.Now()
	mustNewUser := func(
		id uint64,
		name string,
		username string,
		email string,
		btcBalance float64,
		usdBalance float64,
		createdAt time.Time,
		updatedAt time.Time,
	) entity.User {
		user, _ := entity.NewUser(
			id,
			name,
			username,
			email,
			btcBalance,
			usdBalance,
			createdAt,
			updatedAt,
		)
		return *user
	}

	users := map[uint64]entity.User{
		1: mustNewUser(
			1,
			"John",
			"Doe",
			"johndoe@mail.com",
			0,
			0,
			now,
			now,
		),
		2: mustNewUser(
			2,
			"Jane",
			"Doe",
			"janedoe@mail.com",
			100,
			100,
			now,
			now,
		),
	}
	return &userRepositories.MockUserRepository{
		SaveFunc: func(user *entity.User) error {
			now := time.Now()
			btc, _ := user.Balance.BTC.ToFloat().Float64()
			usd, _ := user.Balance.USD.ToFloat().Float64()
			_, err := entity.NewUser(
				user.ID,
				user.Name,
				user.Username,
				user.Email.Address,
				btc,
				usd,
				now,
				now,
			)
			return err
		},
		DeleteFunc: func(id uint64) error {
			if _, ok := users[id]; !ok {
				return userRepositories.ErrUserNotFound
			}
			return nil
		},
		GetFunc: func(id uint64) (*entity.User, error) {
			user, ok := users[id]
			if !ok {
				return nil, userRepositories.ErrUserNotFound
			}
			return &user, nil
		},
	}
}

func NewMockBitcoinRepository() bitcoinEntity.BTCRepository {
	return &bitcoinRepositories.MockBTCRepository{
		GetPriceFunc: func() bitcoinEntity.BTCPrice {
			return bitcoinEntity.NewBTCPrice(bitcoinEntity.MustNewUSD(100), time.Now())
		},
		SetPriceFunc: func(price bitcoinEntity.USD) error {
			return nil
		},
	}
}
