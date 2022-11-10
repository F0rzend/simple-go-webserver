package main

import (
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
	"github.com/F0rzend/simple-go-webserver/app/server"
)

func main() {
	logger := log.
		Output(zerolog.ConsoleWriter{Out: os.Stderr}).
		With().Caller().
		Logger()

	address := getEnv("ADDRESS", ":8080")
	logger.Info().Msgf("starting endpoints on %s", address)

	userRepository := userrepositories.NewMemoryUserRepository()
	bitcoinRepository := bitcoinrepositories.NewMemoryBTCRepository()

	apiServer := server.NewServer(userRepository, bitcoinRepository, bitcoinRepository)

	srv := &http.Server{
		Addr:              address,
		Handler:           apiServer.GetHTTPHandler(&logger),
		ReadHeaderTimeout: 1 * time.Second,
	}
	log.Error().Err(srv.ListenAndServe()).Send()
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
