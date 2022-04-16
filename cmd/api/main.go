package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	userRepositories "github.com/F0rzend/SimpleGoWebserver/internal/adapters/repositories/user"
	"github.com/F0rzend/SimpleGoWebserver/internal/application"
	server "github.com/F0rzend/SimpleGoWebserver/internal/ports/http"
)

func main() {
	log.Logger = log.
		With().Caller().Logger().
		Output(zerolog.ConsoleWriter{Out: os.Stderr})

	userRepository := userRepositories.NewMemoryUserRepository()
	app := application.NewApplication(userRepository)
	httpServer := server.NewServer(app)

	log.Error().Err(http.ListenAndServe("localhost:8080", httpServer.GetRouter()))
}
