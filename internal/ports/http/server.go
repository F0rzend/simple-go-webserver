package http

import (
	"github.com/go-chi/chi/v5/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/F0rzend/SimpleGoWebserver/internal/application"
	"github.com/F0rzend/SimpleGoWebserver/internal/ports/http/types"
)

type Server struct {
	app       *application.Application
	assembler *types.Assembler
}

func NewServer(app *application.Application) *Server {
	return &Server{
		app:       app,
		assembler: types.NewAssembler(),
	}
}

func (s *Server) GetRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.AllowContentType("application/json"),
	)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", s.CreateUser)
		r.Get("/{id}", s.GetUser)
		r.Put("/{id}", s.UpdateUser)
	})

	return r
}
