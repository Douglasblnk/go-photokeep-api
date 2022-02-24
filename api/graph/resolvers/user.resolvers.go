package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-dynamodb/api/graph/generated"
	"go-dynamodb/api/graph/model"
	"go-dynamodb/internal/dto"
	"go-dynamodb/internal/models"
	userRepository "go-dynamodb/internal/repository/user"
)

func (r *mutationResolver) CreateUser(ctx context.Context, data model.CreateUserInput) (*models.User, error) {
	user := &dto.CreateUserDTO{
		ID:       data.ID,
		Name:     data.Name,
		Password: data.Password,
	}

	return userRepository.CreateUser(user)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id *string, data *model.UpdateUserInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUser(ctx context.Context, id *string) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
