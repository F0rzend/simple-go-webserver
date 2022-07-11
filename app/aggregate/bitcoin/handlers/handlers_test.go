package handlers

//
//import (
//	"net/http"
//	"testing"
//
//	"github.com/go-chi/chi/v5"
//	"github.com/go-chi/chi/v5/middleware"
//
//	bitcoinService "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
//	"github.com/F0rzend/simple-go-webserver/app/tests"
//)
//
//// Tests BitcoinHTTPHandlers.SetBTCPrice (POST /bitcoin endpoint)
//func TestServer_SetBTCPrice(t *testing.T) {
//	t.Parallel()
//
//	handler := getHTTPHandler(t)
//
//	testCases := []struct {
//		name     string
//		body     SetBTCPriceRequest
//		expected int
//	}{
//		{
//			name: "success",
//			body: SetBTCPriceRequest{
//				Price: 1,
//			},
//			expected: http.StatusNoContent,
//		},
//		{
//			name: "invalid amount",
//			body: SetBTCPriceRequest{
//				Price: -1,
//			},
//			expected: http.StatusBadRequest,
//		},
//		{
//			name:     "empty body",
//			body:     SetBTCPriceRequest{},
//			expected: http.StatusBadRequest,
//		},
//	}
//
//	for _, tc := range testCases {
//		tc := tc
//		t.Run(tc.name, func(t *testing.T) {
//			t.Parallel()
//
//			tests.AssertStatus(t, handler, http.MethodPut, "/bitcoin", tc.body, tc.expected)
//		})
//	}
//}
//
//func getHTTPHandler(t *testing.T) http.Handler {
//	t.Helper()
//
//	handlers := NewBitcoinHTTPHandlers(
//		bitcoinService.NewBitcoinService(tests.NewMockBitcoinRepository()),
//	)
//
//	r := chi.NewRouter()
//
//	r.Use(
//		middleware.Logger,
//		middleware.Recoverer,
//		middleware.AllowContentType("application/json"),
//	)
//
//	r.Route("/bitcoin", func(r chi.Router) {
//		r.Get("/", handlers.GetBTCPrice)
//		r.Put("/", handlers.SetBTCPrice)
//	})
//
//	return r
//}
