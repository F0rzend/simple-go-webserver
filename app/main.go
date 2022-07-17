package main

import (
	"net/http"
	"os"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	bitcoinRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
	userRepositories "github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/F0rzend/simple-go-webserver/app/server"
)

func main() {
	logger := log.
		Output(zerolog.ConsoleWriter{Out: os.Stderr}).
		With().Caller().
		Logger()

	address := getEnv("ADDRESS", ":8080")
	logger.Info().Msgf("starting endpoints on %s", address)

	userRepository := userRepositories.NewMemoryUserRepository()
	bitcoinRepository, err := bitcoinRepositories.NewMemoryBTCRepository(bitcoinEntity.MustNewUSD(100))
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	apiServer := server.NewServer(userRepository, bitcoinRepository)

	if err := http.ListenAndServe(
		address,
		apiServer.GetHTTPHandler(logger),
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
