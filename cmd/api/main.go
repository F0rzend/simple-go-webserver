package main

import (
	"github.com/F0rzend/SimpleGoWebserver/internal/domain"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	btcRepositories "github.com/F0rzend/SimpleGoWebserver/internal/adapters/repositories/btc"
	userRepositories "github.com/F0rzend/SimpleGoWebserver/internal/adapters/repositories/user"
	"github.com/F0rzend/SimpleGoWebserver/internal/application"
	server "github.com/F0rzend/SimpleGoWebserver/internal/ports/http"
)

func main() {
	log.Logger = log.
		With().Caller().Logger().
		Output(zerolog.ConsoleWriter{Out: os.Stderr})

	userRepository := userRepositories.NewMemoryUserRepository()
	btcRepository := btcRepositories.NewMemoryBTCRepository(domain.USDFromFloat(100))
	app := application.NewApplication(userRepository, btcRepository)
	httpServer := server.NewServer(app)

	log.Error().Err(http.ListenAndServe("localhost:8080", httpServer.GetRouter()))
}
