package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/zraney/quotes-starter/gqlgen/graph/generated"
	"github.com/zraney/quotes-starter/gqlgen/graph/model"
)

// NewQuote is the resolver for the newQuote field.
func (r *mutationResolver) NewQuote(ctx context.Context, input model.QuoteInput) (*model.Response, error) {
	quote := &model.Quote{
		Quote:  input.Quote,
		Author: input.Author,
	}

	marshalledData, err := json.Marshal(quote)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(marshalledData)

	request, requestErr := http.NewRequest("POST", "http://34.149.8.254/quotes/", b)
	request.Header.Set("x-api-key", "COCKTAILSAUCE")

	if requestErr != nil {
		return nil, requestErr
	}

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	responseData, responseDataErr := ioutil.ReadAll(resp.Body)
	if responseDataErr != nil {
		return nil, responseDataErr
	}

	var responseObject model.Response
	json.Unmarshal(responseData, &responseObject)

	return &responseObject, nil
}

// DeleteQuote is the resolver for the deleteQuote field.
func (r *mutationResolver) DeleteQuote(ctx context.Context, id string) (*string, error) {

	requestUrl := "http://34.149.8.254/quotes/" + id
	request, err := http.NewRequest("DELETE", requestUrl, nil)
	request.Header.Set("x-api-key", "COCKTAILSAUCE")

	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(request)

	return &resp.Status, nil
}

// RandomQuote is the resolver for the randomQuote field.
func (r *queryResolver) RandomQuote(ctx context.Context) (*model.Quote, error) {
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

	var newQuote model.Quote
	json.Unmarshal(responseData, &newQuote)

	return &newQuote, nil
}

// QuoteByID is the resolver for the quoteByID field.
func (r *queryResolver) QuoteByID(ctx context.Context, id string) (*model.Quote, error) {
	requestUrl := "http://34.149.8.254/quotes/" + id
	request, err := http.NewRequest("GET", requestUrl, nil)
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
