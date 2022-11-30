package queries

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/ducka/demoroo/internal/core/cqrs"
	"github.com/ducka/demoroo/internal/domain/dtos"
	"github.com/ducka/demoroo/internal/domain/entities"
)

type BranchSearchQueryHandler struct {
	repository    cqrs.Repository[entities.Branch]
	altTermsQuery cqrs.QueryDispatcher[AlternateSearchTermsQuery, AlternateSearchTermsQueryResult]
}

func NewBranchSearchQueryHandler(
	repository cqrs.Repository[entities.Branch],
	altTermsQuery cqrs.QueryDispatcher[AlternateSearchTermsQuery, AlternateSearchTermsQueryResult]) BranchSearchQueryHandler {
	return BranchSearchQueryHandler{
		repository:    repository,
		altTermsQuery: altTermsQuery,
	}
}

func (t BranchSearchQueryHandler) Handle(ctx context.Context, query BranchSearchQuery) (BranchSearchQueryResult, error) {

	// Query handlers are meant to be close to the metal, so it's ok for them to contain knowledge about the storage
	// technology they're querying. As such, I'd expect a query handler to contain:
	//
	// 1) References to clients for storage technologies such as postgres, elastic search, etc.
	// 2) References to external services that may provide data for enriching the query results
	// 3) A reference to the query dispatcher if there's a need to compose data from multiple sub queries
	// 4) I would not normally expect repositories to be used with query handlers. Repositories are normally
	//    used on the command side of the system when you're dealing with entities (as opposed to materialised
	//    views of data - i.e. query results)
	//
	// As for elastic search, I'd like to see us use a DSL for the composition of elastic search queries:
	/*
		// connect to an ElasticSearch instance
		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			 log.Fatalf("Failed creating client: %s", err)
		}

		// run a boolean search query
		res, err := esquery.Search().
			 Query(
			 	esquery.
					Bool().
					Must(esquery.Term("title", "Go and Stuff")).
					Filter(esquery.Term("tag", "tech")),
			).
			Aggs(
				esquery.Avg("average_score", "score"),
				esquery.Max("max_score", "score"),
			).
			Size(20).
			Run(
				es,
				es.Search.WithContext(context.TODO()),
				es.Search.WithIndex("test"),
			)
		if err != nil {
			 log.Fatalf("Failed searching for stuff: %s", err)
		}

		defer res.Body.Close()
	*/

	// First perform a direct search for branches based off the user's specified search term
	branches := t.termSearch(query.SearchText, false)

	// Retrieve all the alternate search terms for the user's specified search term
	var altTerms, _ = t.altTermsQuery.Execute(ctx, AlternateSearchTermsQuery{
		SearchText: query.SearchText,
	})

	// Search for all the branches that match the alternate search terms
	for _, v := range altTerms.Terms {
		branches = append(branches, t.termSearch(v, true)...)
	}

	// simulate latency
	sleepFor := rand.Intn(500)
	time.Sleep(time.Duration(sleepFor) * time.Millisecond)

	return BranchSearchQueryResult{
		TotalCount: len(branches),
		Results:    branches,
	}, nil
}

func (t BranchSearchQueryHandler) termSearch(term string, isAlt bool) []dtos.BranchSearchResultItem {
	results := []dtos.BranchSearchResultItem{}

	for _, v := range t.repository.GetAll() {
		if strings.HasPrefix(strings.ToLower(v.Name), strings.ToLower(term)) {
			results = append(results, mapBranchSearchQueryResultItem(v, isAlt))
		}
	}

	return results
}

// Regarding mapping functions...
//
// Some thought should be given to how we can better organise our mapping functions, as you may end up
// with quite a large number of them give the amount of mapping that can occur in n-layered architectures.
func mapBranchSearchQueryResultItem(branch entities.Branch, isAlt bool) dtos.BranchSearchResultItem {
	return dtos.BranchSearchResultItem{
		ID:   branch.ID,
		Name: branch.Name,
		Alt:  isAlt,
	}
}

type BranchSearchQuery struct {
	SearchText string
}

type BranchSearchQueryResult struct {
	TotalCount int
	Results    []dtos.BranchSearchResultItem
}
