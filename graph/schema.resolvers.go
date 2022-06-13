package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/YutaUra/go-graphql-server/ent"
	"github.com/YutaUra/go-graphql-server/graph/generated"
	"github.com/YutaUra/go-graphql-server/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*ent.Todo, error) {
	todo, err := r.client.Todo.Create().SetText(input.Text).SetIsDone(false).SetOwnerID(input.UserID).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create Todo failed.\n%v", err)
	}
	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodoInput) (*ent.Todo, error) {
	statement := r.client.Todo.UpdateOneID(input.ID)
	if input.IsDone != nil {
		statement = statement.SetIsDone(*input.IsDone)
	}
	if input.Text != nil {
		statement = statement.SetText(*input.Text)
	}
	todo, err := statement.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("update Todo failed.\n%v", err)
	}
	return todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodoInput) (*ent.Todo, error) {
	todo, err := r.client.Todo.Get(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("delete Todo failed.\n%v", err)
	}
	err = r.client.Todo.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("delete Todo failed.\n%v", err)
	}
	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
	user, err := r.client.User.Create().SetName(input.Name).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create User failed.\n%v", err)
	}
	return user, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*ent.Todo, error) {
	todos, err := r.client.Todo.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("Todos failed.\n%v", err)
	}
	return todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *ent.Todo) (*ent.User, error) {
	user, err := obj.QueryOwner().First(ctx)
	if err != nil {
		if err.Error() == "ent: user not found" {
			return nil, nil
		}
		return nil, fmt.Errorf("query User failed.\n%v", err)
	}
	return user, nil
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *todoResolver) Done(ctx context.Context, obj *ent.Todo) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
