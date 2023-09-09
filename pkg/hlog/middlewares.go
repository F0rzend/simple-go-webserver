package hlog

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/rs/xid"
)

type (
	HTTPMiddleware = func(next http.Handler) http.Handler
)

func LoggerInjectionMiddleware(logger *slog.Logger) HTTPMiddleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := ContextWithLogger(r.Context(), logger)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func RequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		logger := GetLoggerFromContext(ctx)

		requestGroup := slog.Group(
			"request",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
			slog.String("remote_addr", r.RemoteAddr),
		)

		start := time.Now()

		ww := wrapWriter(w)
		next.ServeHTTP(ww, r)

		responseGroup := slog.Group(
			"response",
			slog.Int("status_code", ww.StatusCode()),
			slog.String("body", string(ww.Body())),
			slog.Duration("duration", time.Since(start)),
		)

		logger.InfoContext(ctx, "request", requestGroup, responseGroup)
	})
}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		logger := GetLoggerFromContext(ctx)

		requestID, ok := GetRequestIDFromContext(ctx)
		if !ok {
			requestID = xid.New()
			ctx = ContextWithRequestID(ctx, requestID)
		}

		logger = logger.With(slog.String("request_id", requestID.String()))
		ctx = ContextWithLogger(ctx, logger)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
