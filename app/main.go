package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"

	bitcoinEndpoints "github.com/F0rzend/simple-go-webserver/app/bitcoin/http/endpoints"
	userEndpoints "github.com/F0rzend/simple-go-webserver/app/user/http/endpoints"

	bitcoinService "github.com/F0rzend/simple-go-webserver/app/bitcoin/service"
	userService "github.com/F0rzend/simple-go-webserver/app/user/service"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.
		With().Caller().Logger().
		Output(zerolog.ConsoleWriter{Out: os.Stderr})

	userHTTPEndpoints := userEndpoints.NewUserHTTPEndpoints(userService.MustUserService())
	bitcoinHTTPEndpoints := bitcoinEndpoints.NewBitcoinHTTPEndpoints(bitcoinService.MustBitcoinService())

	router := getRouter(
		userHTTPEndpoints,
		bitcoinHTTPEndpoints,
	)

	address := getEnv("ADDRESS", ":8080")
	log.Info().Msgf("starting endpoints on %s", address)

	if err := http.ListenAndServe(
		address,
		router,
	); err != nil {
		log.Error().Err(err).Send()
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getRouter(
	userRoutes *userEndpoints.UserHTTPEndpoints,
	bitcoinRoutes *bitcoinEndpoints.BitcoinHTTPEndpoints,
) http.Handler {
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.AllowContentType("application/json"),
	)

	userRoutes.Register(r)
	bitcoinRoutes.Register(r)

	return r
}
