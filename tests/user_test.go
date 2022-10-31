package tests

import (
	"fmt"
	"net/http"
)

const (
	btcBalanceKey = "btc_balance"
	usdBalanceKey = "usd_balance"
)

func (s *TestSuite) TestUser() {
	user := JSON{
		"name":     "John Doe",
		"username": "john1234",
		"email":    "john1234@mail.com",
	}

	var userLocation string

	s.Run("create user", func() {
		userLocation = s.e().POST("/users").
			WithJSON(user).
			Expect().
			Status(http.StatusCreated).
			ContentType(ContentType, Encoding).
			Header(LocationHeader).Raw()
	})

	s.Run("get created user", func() {
		s.e().GET(userLocation).
			Expect().
			Status(http.StatusOK).
			ContentType(ContentType, Encoding).
			JSON().Object().ContainsMap(user)
	})

	s.Run("check balance", func() {
		s.e().GET(userLocation).
			Expect().
			Status(http.StatusOK).
			ContentType(ContentType, Encoding).
			JSON().Object().
			ValueEqual("btc_balance", "0").
			ValueEqual("usd_balance", "0")
	})
}

func (s *TestSuite) TestUserUSDManipulation() {
	user := JSON{
		"name":     "John Doe",
		"username": "john1234",
		"email":    "john1234@mail.com",
	}

	userLocation := s.e().POST("/users").
		WithJSON(user).
		Expect().
		Status(http.StatusCreated).
		ContentType(ContentType, Encoding).
		Header(LocationHeader).Raw()

	path := fmt.Sprintf("%s/%s", userLocation, "usd")

	s.Run("deposit 100 USD", func() {
		s.e().POST(path).
			WithJSON(JSON{
				"action": "deposit",
				"amount": 100,
			}).
			Expect().
			Status(http.StatusNoContent).
			ContentType(ContentType, Encoding)

		s.e().GET(userLocation).
			Expect().
			Status(http.StatusOK).
			ContentType(ContentType, Encoding).
			JSON().Object().
			ValueEqual(btcBalanceKey, "0").
			ValueEqual(usdBalanceKey, "100")
	})

	s.Run("withdraw 50 USD", func() {
		s.e().POST(path).
			WithJSON(JSON{
				"action": "withdraw",
				"amount": 50,
			}).
			Expect().
			Status(http.StatusNoContent).
			ContentType(ContentType, Encoding)

		s.e().GET(userLocation).
			Expect().
			Status(http.StatusOK).
			ContentType(ContentType, Encoding).
			JSON().Object().
			ValueEqual(btcBalanceKey, "0").
			ValueEqual(usdBalanceKey, "50")
	})
}

func (s *TestSuite) TestUserBTCManipulation() {
	user := JSON{
		"name":     "John Doe",
		"username": "john1234",
		"email":    "john1234@mail.com",
	}

	userLocation := s.e().POST("/users").
		WithJSON(user).
		Expect().
		Status(http.StatusCreated).
		ContentType(ContentType, Encoding).
		Header(LocationHeader).Raw()

	usdPath := fmt.Sprintf("%s/%s", userLocation, "usd")
	btcPath := fmt.Sprintf("%s/%s", userLocation, "btc")

	s.Run("buy 1 BTC", func() {
		s.e().POST(usdPath).
			WithJSON(JSON{
				"action": "deposit",
				"amount": 100,
			}).
			Expect().
			Status(http.StatusNoContent).
			ContentType(ContentType, Encoding)

		s.e().POST(btcPath).
			WithJSON(JSON{
				"action": "buy",
				"amount": 1,
			}).
			Expect().
			Status(http.StatusNoContent).
			ContentType(ContentType, Encoding)

		s.e().GET(userLocation).
			Expect().
			Status(http.StatusOK).
			ContentType(ContentType, Encoding).
			JSON().Object().
			ValueEqual(btcBalanceKey, "1").
			ValueEqual(usdBalanceKey, "0")
	})

	s.Run("sell 1 BTC", func() {
		s.e().POST(btcPath).
			WithJSON(JSON{
				"action": "sell",
				"amount": 1,
			}).
			Expect().
			Status(http.StatusNoContent).
			ContentType(ContentType, Encoding)

		s.e().GET(userLocation).
			Expect().
			Status(http.StatusOK).
			ContentType(ContentType, Encoding).
			JSON().Object().
			ValueEqual(btcBalanceKey, "0").
			ValueEqual(usdBalanceKey, "100")
	})
}
