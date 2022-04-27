package main

import (
	"net/http"
	"os"

	"github.com/F0rzend/SimpleGoWebserver/internal/domain"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	btcRepositories "github.com/F0rzend/SimpleGoWebserver/internal/adapters/repositories/btc"
	userRepositories "github.com/F0rzend/SimpleGoWebserver/internal/adapters/repositories/user"
	"github.com/F0rzend/SimpleGoWebserver/internal/application"
	server "github.com/F0rzend/SimpleGoWebserver/internal/ports/http"
)

var DefaultBitcoinPrice = domain.USDFromCent(100_000_000) // nolint: gomnd

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

	log.Error().Err(http.ListenAndServe("localhost:8080", httpServer.GetRouter()))
}
