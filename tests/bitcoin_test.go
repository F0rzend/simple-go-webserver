package tests

import (
	"net/http"
)

func (s *TestSuite) TestBitcoinPriceManipulations() {
	const endpoint = "/bitcoin"

	s.Run("default price", func() {
		s.e().GET(endpoint).
			Expect().
			Status(http.StatusOK).
			JSON().Object().
			Value("btc").Object().ValueEqual("price", "100")
	})

	s.Run("set price", func() {
		s.e().PUT(endpoint).
			WithJSON(JSON{
				"price": 1,
			}).
			Expect().
			Status(http.StatusNoContent).
			ContentType(ContentType, Encoding)
	})

	s.Run("check changed price", func() {
		s.e().GET(endpoint).
			Expect().
			Status(http.StatusOK).
			JSON().Object().
			Value("btc").Object().ValueEqual("price", "1")
	})

	s.Run("return the default price", func() {
		s.e().PUT(endpoint).
			WithJSON(JSON{
				"price": 100,
			}).
			Expect().
			Status(http.StatusNoContent).
			ContentType(ContentType, Encoding)
	})
}
