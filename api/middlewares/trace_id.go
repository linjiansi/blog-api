package middlewares

import (
	"context"
	"sync"
)

var (
	logNo = 1
	mu    sync.Mutex
)

type traceIDKey struct{}

func newTraceID() int {
	var no int

	mu.Lock()
	no = logNo
	logNo += 1
	mu.Unlock()

	return no
}

func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func GetTraceID(ctx context.Context) int {
	anyValue := ctx.Value(traceIDKey{})

	if id, ok := anyValue.(int); ok {
		return id
	} else {
		return 0
	}
}
