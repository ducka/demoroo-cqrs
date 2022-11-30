package logging

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func RegisterLogger() fx.Option {
	return fx.Provide(
		func(lifecycle fx.Lifecycle) *zap.Logger {
			logger, _ := zap.NewProduction()

			lifecycle.Append(fx.Hook{
				OnStop: func(ctx context.Context) error {
					logger.Sync()
					return nil
				},
			})

			return logger
		},
	)
}
