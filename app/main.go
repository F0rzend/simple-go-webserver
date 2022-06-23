package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/F0rzend/simple-go-webserver/app/application"
	"github.com/F0rzend/simple-go-webserver/app/ports/http/server"
)

func main() {
	log.Logger = log.
		With().Caller().Logger().
		Output(zerolog.ConsoleWriter{Out: os.Stderr})

	app, err := application.NewApplication()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	httpServer := server.NewServer(app)

	address := getEnv("ADDRESS", ":8080")
	log.Info().Msgf("starting server on %s", address)

	if err := http.ListenAndServe(
		address,
		httpServer.GetRouter(),
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
