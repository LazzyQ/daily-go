package addservice

import (
	"context"
	"github.com/go-kit/kit/log"
)

type Middleware func(Service) Service

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (l loggingMiddleware) Sum(ctx context.Context, a, b int) (int, error) {
	panic("implement me")
}

func (l loggingMiddleware) Concat(ctx context.Context, a, b string) (string, error) {
	panic("implement me")
}

func LogginMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{logger, next}
	}
}
