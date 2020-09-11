package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/JonMallozzi/Go_Fiber_GraphQL_DEMO/graph/generated"
	"github.com/JonMallozzi/Go_Fiber_GraphQL_DEMO/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.UsersRepo.CreateUser(input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, username *string, password *string, email *string, dateOfBirth *time.Time) (*model.User, error) {
	return r.UsersRepo.UpdateUser(id, username, password, email, dateOfBirth)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (string, error) {
	return r.UsersRepo.DeleteUser(id)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.UsersRepo.GetUsers()
}

func (r *queryResolver) UserByID(ctx context.Context, id string) (*model.User, error) {
	return r.UsersRepo.GetUserByID(id)
}

func (r *queryResolver) UserByUsername(ctx context.Context, username string) (*model.User, error) {
	return r.UsersRepo.GetUserByUsername(username)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
