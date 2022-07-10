package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	bitcoinHTTPHandlers "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/handlers"
	bitcoinRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
	bitcoinService "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"

	userHTTPHandlers "github.com/F0rzend/simple-go-webserver/app/aggregate/user/handlers"
	userRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
	userService "github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
)

type Server struct {
	userRoutes    *userHTTPHandlers.UserHTTPHandlers
	bitcoinRoutes *bitcoinHTTPHandlers.BitcoinHTTPHandlers
}

func NewServer() (*Server, error) {
	bitcoinRepository, err := bitcoinRepositories.NewMemoryBTCRepository(bitcoinEntity.MustNewUSD(100))
	if err != nil {
		return nil, err
	}
	userRepository := userRepositories.NewMemoryUserRepository()

	bitcoinRoutes := bitcoinHTTPHandlers.NewBitcoinHTTPHandlers(bitcoinService.NewBitcoinService(bitcoinRepository))
	userRoutes := userHTTPHandlers.NewUserHTTPHandlers(userService.NewUserService(userRepository, bitcoinRepository))

	return &Server{
		userRoutes:    userRoutes,
		bitcoinRoutes: bitcoinRoutes,
	}, nil
}

func (s *Server) GetHTTPHandler() http.Handler {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.AllowContentType("application/json"),
	)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", s.userRoutes.CreateUser)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.userRoutes.GetUser)
			r.Put("/", s.userRoutes.UpdateUser)
			r.Get("/balance", s.userRoutes.GetUserBalance)

			r.Post("/bitcoin", s.userRoutes.ChangeBTCBalance)
			r.Post("/usd", s.userRoutes.ChangeUSDBalance)
		})
	})
	r.Route("/bitcoin", func(r chi.Router) {
		r.Get("/", s.bitcoinRoutes.GetBTCPrice)
		r.Put("/", s.bitcoinRoutes.SetBTCPrice)
	})

	return r
}
