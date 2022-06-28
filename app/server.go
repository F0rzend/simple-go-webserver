package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	bitcoinHTTPHandlers "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/handlers"
	bitcoinService "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"
	userHTTPHandlers "github.com/F0rzend/simple-go-webserver/app/aggregate/user/handlers"
	userService "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
)

type Server struct {
	userRoutes    *userHTTPHandlers.UserHTTPHandlers
	bitcoinRoutes *bitcoinHTTPHandlers.BitcoinHTTPHandlers
}

func NewServer() *Server {
	userRoutes := userHTTPHandlers.NewUserHTTPHandlers(userService.MustUserService())
	bitcoinRoutes := bitcoinHTTPHandlers.NewBitcoinHTTPHandlers(bitcoinService.MustBitcoinService())

	return &Server{
		userRoutes:    userRoutes,
		bitcoinRoutes: bitcoinRoutes,
	}
}

func (s *Server) GetHTTPHandler() http.Handler {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.AllowContentType("application/json"),
	)

	s.userRoutes.SetRoutes(r)
	s.bitcoinRoutes.SetRoutes(r)

	return r
}
