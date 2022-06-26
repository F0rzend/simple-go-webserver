package endpoints

import (
	"github.com/F0rzend/simple-go-webserver/app/user/service"
	"github.com/go-chi/chi/v5"
)

type UserHTTPEndpoints struct {
	service *service.UserService
}

func NewUserHTTPEndpoints(service *service.UserService) *UserHTTPEndpoints {
	return &UserHTTPEndpoints{
		service: service,
	}
}

func (u *UserHTTPEndpoints) Register(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", u.CreateUser)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", u.GetUser)
			r.Put("/", u.UpdateUser)
			r.Get("/balance", u.GetUserBalance)

			r.Post("/bitcoin", u.ChangeBTCBalance)
			r.Post("/usd", u.ChangeUSDBalance)
		})
	})
}
