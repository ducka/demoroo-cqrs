package queries

import (
	"context"
	"math/rand"
	"strings"
	"time"
)

type AlternateSearchTermsQueryHandler struct {
}

func NewAlternateSearchTermsQueryHandler() AlternateSearchTermsQueryHandler {
	return AlternateSearchTermsQueryHandler{}
}

func (t AlternateSearchTermsQueryHandler) Handle(ctx context.Context, query AlternateSearchTermsQuery) (AlternateSearchTermsQueryResult, error) {

	// Simulating alternative search terms results
	terms := []string{}
	switch strings.ToLower(query.SearchText) {
	case "kfc":
		terms = []string{"mc donalds"}
	case "waitroes":
		terms = []string{"tesco"}
	}

	// simulate latency
	sleepFor := rand.Intn(200)
	time.Sleep(time.Duration(sleepFor) * time.Millisecond)

	return AlternateSearchTermsQueryResult{
		Terms: terms,
	}, nil
}

type AlternateSearchTermsQuery struct {
	SearchText string
}

type AlternateSearchTermsQueryResult struct {
	Terms []string
}
