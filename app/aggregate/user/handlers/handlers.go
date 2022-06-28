package handlers

import (
	"github.com/go-chi/chi/v5"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/service"
)

type UserHTTPHandlers struct {
	service *service.UserService
}

func NewUserHTTPHandlers(userService *service.UserService) *UserHTTPHandlers {
	return &UserHTTPHandlers{
		service: userService,
	}
}

func (h *UserHTTPHandlers) SetRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.createUser)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.getUser)
			r.Put("/", h.updateUser)
			r.Get("/balance", h.getUserBalance)

			r.Post("/bitcoin", h.changeBTCBalance)
			r.Post("/usd", h.changeUSDBalance)
		})
	})
}
