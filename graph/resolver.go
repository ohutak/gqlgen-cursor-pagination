package graph

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/kabece/gqlgen-chatroom/graph/generated"
	"github.com/kabece/gqlgen-chatroom/graph/model"
)

type Resolver struct {
	Results  map[string]model.SearchResultListings
	Listings map[string][]model.Listing
}

func NewResolver() generated.Config {
	const nResults = 20
	const nListingsPerSearch = 100
	r := Resolver{}
	r.Results = make(map[string]model.SearchResultListings, nResults)
	r.Listings = make(map[string][]model.Listing, nResults)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nResults; i++ {
		id := strconv.Itoa(i + 1)
		searchResultsParams := model.SearchResultListings{
			ID: id,
		}
		r.Results[id] = searchResultsParams
		r.Listings[id] = make([]model.Listing, nListingsPerSearch)

		// Generate listings for the Search
		for k := 0; k < nListingsPerSearch; k++ {
			text := fmt.Sprintf("Listing %d", k + 1)

			mockListing := model.Listing{
				ID: strconv.Itoa(k + 1),
				Title: &text,
			}

			r.Listings[id][k] = mockListing
		}
	}

	return generated.Config{
		Resolvers: &r,
	}
}
