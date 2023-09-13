package hlog

import (
	"context"

	"github.com/rs/xid"
)

type requestIDKey struct{}

func ContextWithRequestID(ctx context.Context, requestID xid.ID) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

func GetRequestIDFromContext(ctx context.Context) (xid.ID, bool) {
	requestID, ok := ctx.Value(requestIDKey{}).(xid.ID)

	return requestID, ok
}
