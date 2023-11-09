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
	Srps  map[string]model.Srp
	Listings map[string][]model.Listing
}

func NewResolver() generated.Config {
	const nSrps = 20
	const nListingsPerSrp = 100
	r := Resolver{}
	r.Srps = make(map[string]model.Srp, nSrps)
	r.Listings = make(map[string][]model.Listing, nSrps)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nSrps; i++ {
		id := strconv.Itoa(i + 1)
		mockChatRoom := model.Srp{
			ID: id,
			Name: fmt.Sprintf("Srp Page %s", id),
		}
		r.Srps[id] = mockChatRoom
		r.Listings[id] = make([]model.Listing, nListingsPerSrp)

		// Generate messages for the ChatRoom
		for k := 0; k < nListingsPerSrp; k++ {
			text := fmt.Sprintf("Listing %d", k + 1)

			mockMessage := model.Listing{
				ID: strconv.Itoa(k + 1),
				Title: &text,
			}

			r.Listings[id][k] = mockMessage
		}
	}

	return generated.Config{
		Resolvers: &r,
	}
}
