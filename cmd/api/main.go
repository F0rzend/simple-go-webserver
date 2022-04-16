package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"

	userRepositories "github.com/F0rzend/SimpleGoWebserver/internal/adapters/repositories/user"
	"github.com/F0rzend/SimpleGoWebserver/internal/application"
	httpServer "github.com/F0rzend/SimpleGoWebserver/internal/ports/http"
)

func main() {
	log.Logger = log.
		With().Caller().Logger().
		Output(zerolog.ConsoleWriter{Out: os.Stderr})

	userRepository := userRepositories.NewMemoryUserRepository()
	app := application.NewApplication(userRepository)
	server := httpServer.NewServer(app)

	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.AllowContentType("application/json"),
	)

	router.Route("/users", func(r chi.Router) {
		r.Post("/", server.CreateUser)
	})

	log.Error().Err(http.ListenAndServe("localhost:8080", router))
}
