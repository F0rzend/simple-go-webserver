package main

import (
	"fmt"
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
		log.Fatal().Err(err).Msg("failed to create btc repository")
	}
	httpServer := server.NewServer(app)

	address := getEnv("ADDRESS", ":8080")
	log.Info().Msg(fmt.Sprintf("starting server on %s", address))
	log.Error().Err(http.ListenAndServe(
		address,
		httpServer.GetRouter(),
	))
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
