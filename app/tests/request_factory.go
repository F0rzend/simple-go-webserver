package tests

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"os"
	"testing"

	"github.com/F0rzend/simple-go-webserver/pkg/hlog"

	"github.com/gavv/httpexpect"
)

type requestFactory struct {
	t *testing.T
}

func newRequestFactoryWithTestLogger(t *testing.T) httpexpect.RequestFactory {
	return &requestFactory{t: t}
}

func (rf *requestFactory) NewRequest(method, target string, body io.Reader) (*http.Request, error) {
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	ctx := hlog.ContextWithLogger(context.Background(), logger)

	return http.NewRequestWithContext(ctx, method, target, body)
}
