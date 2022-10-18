package common

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

type ApplicationError struct {
	HTTPStatus   int    `json:"-"`
	ErrorMessage string `json:"error"`
}

func (e ApplicationError) Error() string {
	return e.ErrorMessage
}

func NewApplicationError(httpStatus int, message string) ApplicationError {
	return ApplicationError{
		HTTPStatus:   httpStatus,
		ErrorMessage: message,
	}
}

func (e ApplicationError) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatus)
	return nil
}

func RenderHTTPError(w http.ResponseWriter, r *http.Request, err error) {
	ctx := r.Context()

	switch err := err.(type) {
	case nil:
	case ApplicationError:
		logError(ctx, render.Render(w, r, err))
		return
	default:
		logError(ctx, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func logError(ctx context.Context, err error) {
	logger := log.Ctx(ctx)
	if err != nil {
		logger.Error().Err(err).Send()
	}
}
