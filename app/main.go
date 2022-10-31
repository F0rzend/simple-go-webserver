package main

import (
	"net/http"
	"os"
	"time"

	bitcoinentity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"

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
	err := bitcoinRepository.SetPrice(getDefaultBTCPrice())
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to set bitcoin price")
	}

	apiServer := server.NewServer(userRepository, bitcoinRepository, bitcoinRepository)

	if err := http.ListenAndServe(
		address,
		apiServer.GetHTTPHandler(&logger),
	); err != nil {
		log.Error().Err(err).Send()
	}
}

func getDefaultBTCPrice() bitcoinentity.BTCPrice {
	price, _ := bitcoinentity.NewBTCPrice(bitcoinentity.NewUSD(100), time.Now())
	return price
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
