package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/zraney/quotes-starter/gqlgen/graph/generated"
	"github.com/zraney/quotes-starter/gqlgen/graph/model"
)

// GetRandomQuote is the resolver for the getRandomQuote field.
func (r *queryResolver) GetRandomQuote(ctx context.Context) (*model.Quote, error) {

	request, err := http.NewRequest("GET", "http://34.149.8.254/quotes/", nil)
	request.Header.Set("x-api-key", "COCKTAILSAUCE")

	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(request)

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseObject model.Quote
	json.Unmarshal(responseData, &responseObject)

	return &responseObject, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
