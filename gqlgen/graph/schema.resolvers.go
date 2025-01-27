package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
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
	key := ctx.Value("x-api-key").(string)
	request.Header.Set("x-api-key", key)

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

	switch resp.StatusCode {
	case 201:
		return &responseObject, nil
	case 400:
		return nil, errors.New("quote and author must be greater than 3 characters")
	case 401:
		return nil, errors.New("Not Authorized")
	}
	return nil, err
}

// DeleteQuote is the resolver for the deleteQuote field.
func (r *mutationResolver) DeleteQuote(ctx context.Context, id string) (*string, error) {
	_, errorResponse := r.Query().QuoteByID(ctx, &id)

	if errorResponse != nil && errorResponse.Error() == "Invalid ID" {
		return nil, errors.New("Invalid ID")
	}
	requestUrl := "http://34.149.8.254/quotes/" + id
	request, err := http.NewRequest("DELETE", requestUrl, nil)
	key := ctx.Value("x-api-key").(string)
	request.Header.Set("x-api-key", key)

	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	return &resp.Status, nil
}

// RandomQuote is the resolver for the randomQuote field.
func (r *queryResolver) RandomQuote(ctx context.Context) (*model.Quote, error) {
	request, err := http.NewRequest("GET", "http://34.149.8.254/quotes/", nil)
	key := ctx.Value("x-api-key").(string)
	request.Header.Set("x-api-key", key)

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

	switch resp.StatusCode {
	case 200:
		return &newQuote, nil
	case 401:
		return nil, errors.New("Not Authorized")
	}
	return nil, err
}

// QuoteByID is the resolver for the quoteByID field.
func (r *queryResolver) QuoteByID(ctx context.Context, id *string) (*model.Quote, error) {
	requestUrl := "http://34.149.8.254/quotes/" + *id
	request, err := http.NewRequest("GET", requestUrl, nil)
	key := ctx.Value("x-api-key").(string)
	request.Header.Set("x-api-key", key)

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

	switch resp.StatusCode {
	case 200:
		return &responseObject, nil
	case 400:
		return nil, errors.New("Invalid ID")
	case 401:
		return nil, errors.New("Not Authorized")
	}
	return nil, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
