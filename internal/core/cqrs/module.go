package cqrs

import (
	"reflect"

	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/fx"
)

func RegisterQueryHandler[TQuery any, TResult any](queryHandlerConstructor interface{}) fx.Option {
	queryName := reflect.TypeOf(new(TQuery)).Name()

	return fx.Module(
		queryName,
		fx.Provide(
			// Register the query handler as an injected dependency
			fx.Annotate(
				queryHandlerConstructor,
				fx.As(new(QueryHandler[TQuery, TResult])),
			),
			// Register the dispatcher for the query handler as an injected dependency
			func(lifecycle fx.Lifecycle) QueryDispatcher[TQuery, TResult] {
				return mediatrQueryDispatcher[TQuery, TResult]{}
			},
		),
		fx.Invoke(func(handler QueryHandler[TQuery, TResult]) {
			// Register the query handler with the Mediatr plugin
			mediatr.RegisterRequestHandler[TQuery, TResult](handler)
		}),
	)
}

type PipelineMiddleware = mediatr.PipelineBehavior
type RequestHandler = mediatr.RequestHandlerFunc

func RegisterPipelineMiddleware(middleware interface{}) fx.Option {
	return fx.Module(
		"pipeline-behavior",
		fx.Provide(
			// Register the query handler as an injected dependency
			fx.Annotate(
				middleware,
				fx.As(new(PipelineMiddleware)),
			),
		),
		fx.Invoke(func(handler PipelineMiddleware) {
			mediatr.RegisterRequestPipelineBehaviors(handler)
		}),
	)
}
