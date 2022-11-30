package api

import (
	"context"

	"github.com/ducka/demoroo/internal/core/cqrs"
	"github.com/ducka/demoroo/internal/domain/queries"
	"github.com/ducka/demoroo/proto"
)

type BranchSearchService struct {
	branchSearchQueryDispatcher cqrs.QueryDispatcher[queries.BranchSearchQuery, queries.BranchSearchQueryResult]
	proto.UnimplementedBranchSearchServiceServer
}

func NewBranchSearchService(branchSearchQueryDispatcher cqrs.QueryDispatcher[queries.BranchSearchQuery, queries.BranchSearchQueryResult]) *BranchSearchService {
	return &BranchSearchService{
		branchSearchQueryDispatcher: branchSearchQueryDispatcher,
	}
}
func (s *BranchSearchService) Search(ctx context.Context, request *proto.BranchSearchRequest) (*proto.BranchSearchResponse, error) {
	// TODO: improve the mapping experience
	query := mapBranchSearchRequest(request)

	queryResult, err := s.branchSearchQueryDispatcher.Execute(ctx, query)

	response := mapBranchSearchResult(queryResult)

	return response, err
}

// Regarding mapping functions...
//
// Some thought should be given to how we can better organise our mapping functions, as you may end up
// with quite a large number of them give the amount of mapping that can occur in n-layered architectures.
func mapBranchSearchRequest(request *proto.BranchSearchRequest) queries.BranchSearchQuery {
	return queries.BranchSearchQuery{
		SearchText: request.SearchText,
	}
}

func mapBranchSearchResult(result queries.BranchSearchQueryResult) *proto.BranchSearchResponse {
	results := make([]*proto.BranchSearchResultItem, 0)

	for _, v := range result.Results {
		results = append(results, &proto.BranchSearchResultItem{
			Id:   v.ID,
			Name: v.Name,
			Alt:  v.Alt,
		})
	}

	return &proto.BranchSearchResponse{
		TotalCount: int32(result.TotalCount),
		Results:    results,
	}
}
