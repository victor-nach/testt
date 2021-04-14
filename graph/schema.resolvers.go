package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"testt/graph/generated"
	"testt/graph/model"
)

func (r *queryResolver) Urlshortener(ctx context.Context, url string) (string, error) {
	urlMap := &model.URL{}

	// check if long url already exists
	urlMap, err := r.Store.GetByUrl(url)

	// if it doesn't exist, save it
	if err != nil {
		urlMap, err = r.Store.Save(&model.URL{URL: url})
		if err != nil {
			return "", err
		}
	}

	shortUrl := fmt.Sprintf("%v/%v", r.Host, urlMap.Code)
	return shortUrl, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }
