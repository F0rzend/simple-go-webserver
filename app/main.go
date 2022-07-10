package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/F0rzend/simple-go-webserver/app/server"
)

func setupLogger() {
	log.Logger = log.
		With().Caller().Logger().
		Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func main() {
	setupLogger()

	address := getEnv("ADDRESS", ":8080")
	log.Info().Msgf("starting endpoints on %s", address)

	server, err := server.NewServer()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	if err := http.ListenAndServe(
		address,
		server.GetHTTPHandler(),
	); err != nil {
		log.Error().Err(err).Send()
	}
}
