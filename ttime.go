package ttime

import (
	"context"
	"testing"
	"time"
)

type nowKey struct{}

// NowFunc is a function whic returns a time.Time
type NowFunc func() time.Time

// Now is an alternative to time.Now() which uses the NowFunc
// provided on the context, given one has been configured.
// Otherwise, it fallsback to time.Now().
func Now(ctx context.Context) time.Time {
	if v := ctx.Value(nowKey{}); v != nil {
		if now, ok := v.(NowFunc); ok {
			return now()
		}
	}

	return time.Now()
}

// WithNow threads the provided NowFunc onto the provided context.
// This functionality is intended for tests only, which is why it requires a
// testing.T pointer in order to work.
func WithNow(ctx context.Context, t *testing.T, now NowFunc) context.Context {
	t.Helper()

	return context.WithValue(ctx, nowKey{}, now)
}
