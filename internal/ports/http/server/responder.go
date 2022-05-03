package server

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

var (
	errNilResponseWriter = errors.New("response writer is nil")
	errNilRequest        = errors.New("request is nil")
)

type Responder struct {
	writer  http.ResponseWriter
	request *http.Request
}

func NewResponder(w http.ResponseWriter, r *http.Request) (*Responder, error) {
	if w == nil {
		return nil, errNilResponseWriter
	}
	if r == nil {
		return nil, errNilRequest
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

func (r *Responder) Response(response any) {
	if response, ok := response.(render.Renderer); ok {
		if err := render.Render(
			r.writer, r.request, response,
		); err != nil {
			log.Error().Err(err).Msg("error on response error")
		}
		return
	}
	render.Respond(r.writer, r.request, response)
}

func (r *Responder) InternalError(err error) {
	log.Error().Err(err).Msg("internal error")
	r.StatusOnly(http.StatusInternalServerError)
}

func (r *Responder) StatusOnly(code int) {
	r.writer.WriteHeader(code)
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
	r.Header("Location", location)
}
