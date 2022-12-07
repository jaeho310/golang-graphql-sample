package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.21 DO NOT EDIT.

import (
	"context"
	"fmt"
	"graphql-sample/generated"
	gql_types1 "graphql-sample/resolver/gql-types"
)

// CreateUser is the resolver for the create_user field.
func (r *mutationResolver) CreateUser(ctx context.Context, input gql_types1.CreateUserInput) (*gql_types1.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - create_user"))
}

// UserList is the resolver for the user_list field.
func (r *queryResolver) UserList(ctx context.Context) (*gql_types1.UserList, error) {
	panic(fmt.Errorf("not implemented: UserList - user_list"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
