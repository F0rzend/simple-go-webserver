// Package common
// Module errors
// RFC-7807 error handling
package common

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/F0rzend/simple-go-webserver/pkg/hlog"

	"github.com/go-chi/render"
)

type ErrorType string

const (
	InternalServerErrorType ErrorType = "InternalServerError"
	ValueErrorType          ErrorType = "ValueError"
	NotFoundErrorType       ErrorType = "NotFoundError"
)

type HTTPError struct {
	Type     ErrorType `json:"type"`
	Status   int       `json:"status"`
	Title    string    `json:"title,omitempty"`
	Detail   string    `json:"detail,omitempty"`
	Instance string    `json:"instance,omitempty"`

	err error
}

func NewInternalServerError(err error) error {
	return &HTTPError{
		Type:   InternalServerErrorType,
		Status: http.StatusInternalServerError,
		Title:  "Error on our side.",
		err:    err,
	}
}

func (e *HTTPError) Error() string {
	err := fmt.Sprintf("%s#%d", e.Type, e.Status)

	if e.Title != "" {
		err = fmt.Sprintf("%s: %s", err, e.Title)
	}

	return err
}

func (e *HTTPError) Unwrap() error {
	return e.err
}

func (e *HTTPError) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Status)

	return nil
}

func (e *HTTPError) GetType() ErrorType {
	return e.Type
}

func (e *HTTPError) GetStatus() int {
	return e.Status
}

func (e *HTTPError) GetTitle() string {
	return e.Title
}

func (e *HTTPError) GetDetail() string {
	return e.Detail
}

func (e *HTTPError) GetInstance() string {
	return e.Instance
}

func ErrorHandler(
	handler func(w http.ResponseWriter, r *http.Request) error,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			renderError(w, r, err)
		}
	}
}

func renderError(w http.ResponseWriter, r *http.Request, err error) {
	var renderer rendererError
	if !errors.As(err, &renderer) {
		logger := hlog.GetLoggerFromContext(r.Context())
		logger.Error(err.Error())

		renderError(w, r, NewInternalServerError(err))
		return
	}

	if renderingError := render.Render(w, r, renderer); renderingError != nil {
		renderError(w, r, NewInternalServerError(renderingError))
	}
}

type rendererError interface {
	error
	Render(http.ResponseWriter, *http.Request) error
}

func NewValidationError(detail string) error {
	return &HTTPError{
		Type:   ValueErrorType,
		Status: http.StatusBadRequest,
		Detail: fmt.Sprintf("Validation Error: %s", detail),
	}
}

func NewNotFoundError(detail string) error {
	return &HTTPError{
		Type:   NotFoundErrorType,
		Status: http.StatusNotFound,
		Detail: fmt.Sprintf("Not Found Error: %s", detail),
	}
}
