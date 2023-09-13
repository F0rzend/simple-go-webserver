package hlog

import (
	"context"
	"log/slog"
)

var _ slog.Handler = (*discardHandler)(nil)

type discardHandler struct{}

func (discardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}

//nolint:hugeParam
func (discardHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (discardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return discardHandler{}
}

func (discardHandler) WithGroup(_ string) slog.Handler {
	return discardHandler{}
}
