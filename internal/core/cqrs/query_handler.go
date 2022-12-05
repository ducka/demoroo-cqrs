package cqrs

import (
	"context"

	"github.com/mehdihadeli/go-mediatr"
)

//type CommandDispatcher[TCommand any] interface {
//	Send(ctx context.Context, query TCommand) (err error)
//}
//
//type EventDispatcher[TEvent any] interface {
//	Publish(ctx context.Context, query TEvent) (err error)
//}

type QueryDispatcher[TQuery any, TResult any] interface {
	Execute(ctx context.Context, query TQuery) (TResult, error)
}

type QueryHandler[TQuery any, TResult any] interface {
	Handle(ctx context.Context, query TQuery) (TResult, error)
}

type mediatrQueryDispatcher[TQuery any, TResult any] struct {
}

func (t mediatrQueryDispatcher[TQuery, TResult]) Execute(ctx context.Context, query TQuery) (TResult, error) {
	return mediatr.Send[TQuery, TResult](ctx, query)
}
