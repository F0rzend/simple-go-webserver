package common

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/F0rzend/simple-go-webserver/pkg/hlog"

	"github.com/go-chi/render"
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
	case *json.UnmarshalTypeError:
		logError(ctx, render.Render(w, r, MarshalError(err)))
		return
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
	logger := hlog.GetLoggerFromContext(ctx)
	logger.Error("error", err)
}

func MarshalError(err *json.UnmarshalTypeError) render.Renderer {
	return NewApplicationError(http.StatusBadRequest, fmt.Sprintf(
		"The '%s' field must be a %s, but got %s",
		err.Field,
		golangTypeToJSONTypes[err.Type.Kind().String()],
		err.Value,
	))
}

var golangTypeToJSONTypes = map[string]string{
	"bool":       "boolean",
	"int":        "number",
	"int8":       "number",
	"int16":      "number",
	"int32":      "number",
	"int64":      "number",
	"uint":       "number",
	"uint8":      "number",
	"uint16":     "number",
	"uint32":     "number",
	"uint64":     "number",
	"float32":    "number",
	"float64":    "number",
	"complex64":  "number",
	"complex128": "number",
	"array":      "array",
	"slice":      "array",
	"map":        "object",
	"struct":     "object",
	"string":     "string",
}
