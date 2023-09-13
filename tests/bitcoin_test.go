package tests

import (
	"net/http"
)

func (s *TestSuite) TestBitcoinPriceManipulations() {
	const endpoint = "/bitcoin"

	s.e().GET(endpoint).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		Value("btc").Object().ValueEqual("price", "100")

	s.e().PUT(endpoint).
		WithJSON(JSON{
			"price": 1,
		}).
		Expect().
		Status(http.StatusNoContent).
		ContentType(ContentType, Encoding)

	s.e().GET(endpoint).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		Value("btc").Object().ValueEqual("price", "1")
}
