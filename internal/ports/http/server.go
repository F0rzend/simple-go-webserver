package http

import (
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
