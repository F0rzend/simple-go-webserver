package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	btcRepositories "github.com/F0rzend/SimpleGoWebserver/internal/adapters/repositories/btc"
	userRepositories "github.com/F0rzend/SimpleGoWebserver/internal/adapters/repositories/user"
	"github.com/F0rzend/SimpleGoWebserver/internal/application"
	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
	"github.com/F0rzend/SimpleGoWebserver/internal/ports/http/server"
)

var DefaultBitcoinPrice = domain.MustNewUSD(100) // nolint: gomnd

func main() {
	log.Logger = log.
		With().Caller().Logger().
		Output(zerolog.ConsoleWriter{Out: os.Stderr})

	userRepository := userRepositories.NewMemoryUserRepository()
	btcRepository, err := btcRepositories.NewMemoryBTCRepository(DefaultBitcoinPrice)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create btc repository")
	}
	app := application.NewApplication(userRepository, btcRepository)
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
