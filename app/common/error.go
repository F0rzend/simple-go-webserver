package common

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

type ApplicationError struct {
	httpStatus   int
	ErrorMessage string `json:"error"`
}

func (e ApplicationError) Error() string {
	return e.ErrorMessage
}

func NewApplicationError(httpStatus int, message string) ApplicationError {
	return ApplicationError{
		httpStatus:   httpStatus,
		ErrorMessage: message,
	}
}

func (e ApplicationError) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.httpStatus)
	return nil
}

func RenderHTTPError(w http.ResponseWriter, r *http.Request, err error) {
	switch err := err.(type) {
	case nil:
	case ApplicationError:
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
