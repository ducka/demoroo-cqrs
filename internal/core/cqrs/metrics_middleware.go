package cqrs

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"go.uber.org/zap"
)

type MetricsMiddleware struct {
	logger *zap.Logger
}

func NewLoggingMiddleware(logger *zap.Logger) MetricsMiddleware {
	return MetricsMiddleware{
		logger: logger,
	}
}

func (t MetricsMiddleware) Handle(ctx context.Context, request interface{}, next RequestHandler) (interface{}, error) {

	// NB: You would also:
	// 1) Create an APM Span in here
	// 2) Capture metadata from the context and set it on the SPAN.
	start := time.Now()
	requestName := reflect.TypeOf(request).Name()

	t.logger.Info(fmt.Sprintf("Executing request %s", requestName))

	defer func() {
		elapsed := time.Now().Sub(start).Milliseconds()

		t.logger.Info(
			fmt.Sprintf("Request %s took %dms to execute", requestName, elapsed),
		)
	}()

	response, err := next()

	return response, err
}
