package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/handlers"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/service"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/entity"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/handlers"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
)

type Server struct {
	userRoutes    *userhandlers.UserHTTPHandlers
	bitcoinRoutes *bitcoinhandlers.BitcoinHTTPHandlers
}

func getUserIDFromURL(r *http.Request) (uint64, error) {
	const userIDURLKey = "id"

	return strconv.ParseUint(chi.URLParam(r, userIDURLKey), 10, 64) //nolint:gomnd
}

func NewServer(
	userRepository userentity.UserRepository,
	bitcoinRepository bitcoinentity.BTCRepository,
) *Server {
	bitcoinRoutes := bitcoinhandlers.NewBitcoinHTTPHandlers(bitcoinservice.NewBitcoinService(bitcoinRepository))
	userRoutes := userhandlers.NewUserHTTPHandlers(
		userservice.NewUserService(userRepository, bitcoinRepository),
		getUserIDFromURL,
	)

	return &Server{
		userRoutes:    userRoutes,
		bitcoinRoutes: bitcoinRoutes,
	}
}

func (s *Server) GetHTTPHandler(
	logger *zerolog.Logger,
) http.Handler {
	r := chi.NewRouter()

	r.Use(
		middleware.Recoverer,
		middleware.AllowContentType("application/json"),

		hlog.NewHandler(*logger),
		hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).Info().
				Str("method", r.Method).
				Stringer("url", r.URL).
				Int("status", status).
				Int("size", size).
				Dur("duration", duration).
				Send()
		}),
		hlog.RemoteAddrHandler("ip"),
		hlog.RequestIDHandler("req_id", "Request-Id"),
	)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", s.userRoutes.CreateUser)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.userRoutes.GetUser)
			r.Put("/", s.userRoutes.UpdateUser)
			r.Get("/balance", s.userRoutes.GetUserBalance)

			r.Post("/usd", s.userRoutes.ChangeUSDBalance)
			r.Post("/bitcoin", s.userRoutes.ChangeBTCBalance)
		})
	})
	r.Route("/bitcoin", func(r chi.Router) {
		r.Get("/", s.bitcoinRoutes.GetBTCPrice)
		r.Put("/", s.bitcoinRoutes.SetBTCPrice)
	})

	return r
}
