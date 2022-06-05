package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/YutaUra/go-graphql-server/graph/generated"
	"github.com/YutaUra/go-graphql-server/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	for _, todo := range r.todos {
		if todo.ID != input.ID {
			continue
		}
		if input.Done != nil {
			todo.Done = *input.Done
		}
		if input.Text != nil {
			todo.Text = *input.Text
		}
		return todo, nil
	}
	return nil, fmt.Errorf("Todo is not found")
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodoInput) (*model.Todo, error) {
	for i, todo := range r.todos {
		if todo.ID != input.ID {
			continue
		}
		r.todos = append(r.todos[:i], r.todos[i+1:]...)
		return todo, nil
	}
	return nil, fmt.Errorf("Todo is not found")
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
