package domain

import (
	"github.com/ducka/demoroo/internal/core/cqrs"
	"github.com/ducka/demoroo/internal/data"
	"github.com/ducka/demoroo/internal/domain/entities"
	"github.com/ducka/demoroo/internal/domain/queries"
	"go.uber.org/fx"
)

// RegisterApplication composes the dependencies to run the application domain
func RegisterApplication() fx.Option {
	return fx.Module(
		"application",

		// Data layer
		fx.Provide(
			fx.Annotate(
				data.NewBranchRepository,
				fx.As(new(cqrs.Repository[entities.Branch])),
			),
		),

		// Query handler middleware
		cqrs.RegisterPipelineMiddleware(cqrs.NewLoggingMiddleware),

		// Query handlers
		cqrs.RegisterQueryHandler[queries.BranchSearchQuery, queries.BranchSearchQueryResult](queries.NewBranchSearchQueryHandler),
		cqrs.RegisterQueryHandler[queries.AlternateSearchTermsQuery, queries.AlternateSearchTermsQueryResult](queries.NewAlternateSearchTermsQueryHandler),
	)
}
