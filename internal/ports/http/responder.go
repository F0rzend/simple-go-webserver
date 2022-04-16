package http

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	"github.com/F0rzend/SimpleGoWebserver/internal/ports/http/types"
)

var (
	errNilResponseWriter = errors.New("response writer is nil")
)

type Responder struct {
	writer  http.ResponseWriter
	request *http.Request
}

func NewResponder(w http.ResponseWriter, r *http.Request) (*Responder, error) {
	if w == nil {
		return nil, errNilResponseWriter
	}

	return &Responder{writer: w, request: r}, nil
}

func MustNewResponder(w http.ResponseWriter, r *http.Request) *Responder {
	resp, err := NewResponder(w, r)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	return resp
}

func (r *Responder) Response(response render.Renderer) {
	if err := render.Render(
		r.writer, r.request, response,
	); err != nil {
		log.Error().Err(err).Msg("error on response error")
	}
}

func SuccessResponse(result any) render.Renderer {
	return types.Response{
		Ok:     true,
		Result: result,
	}
}

func Error(code int, err error) render.Renderer {
	return types.Response{
		Ok: false,
		Error: &types.HttpError{
			Code:        code,
			Error:       http.StatusText(code),
			Description: err.Error(),
		},
	}
}

func (r *Responder) InternalError() {
	r.Status(http.StatusInternalServerError)
	r.Response(types.InternalError)
}

func (r *Responder) Status(code int) {
	render.Status(r.request, code)
}

func (r *Responder) Header(key, value string) {
	r.writer.Header().Set(key, value)
}

func (r *Responder) ContentType(contentType string) {
	r.Header("Content-Type", contentType)
}

func (r *Responder) LocationHeader(location string) {
	r.writer.Header().Set("Location", location)
}
