package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"github.com/santoshdawanse/graphql-app/graph"
	"github.com/santoshdawanse/graphql-app/graph/model"
	"github.com/sirupsen/logrus"
)

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	logrus.Info("Querying user...")
	return &model.User{}, nil
}

// Todo returns graph.TodoResolver implementation.
func (r *Resolver) Todo() graph.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
