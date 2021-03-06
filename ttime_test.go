package ttime

import (
	"context"
	"fmt"
	"testing"
	"time"
)

var t = &testing.T{}

func Example() {
	var (
		ctx = context.Background()
		fn  = func(ctx context.Context) {
			fmt.Printf("The time is %q", Now(ctx))
		}
	)

	fn(ctx)

	fn(WithNow(ctx, t, func() time.Time {
		return time.Date(2020, 3, 16, 16, 55, 0, 0, time.UTC)
	}))
}
