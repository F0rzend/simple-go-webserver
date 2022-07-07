package common

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

type ServiceError struct {
	HTTPStatus   int    `json:"-"`
	ErrorMessage string `json:"error"`
}

func NewServiceError(status int, message string) ServiceError {
	return ServiceError{HTTPStatus: status, ErrorMessage: message}
}

func (e ServiceError) Error() string {
	return e.ErrorMessage
}

func (e ServiceError) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatus)
	return nil
}

func RenderHTTPError(w http.ResponseWriter, r *http.Request, err error) {
	switch err := err.(type) {
	case nil:
	case ServiceError:
		logError(render.Render(w, r, err))
		return
	default:
		logError(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func logError(err error) {
	if err != nil {
		log.Error().Err(err).Send()
	}
}
