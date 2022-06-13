package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/YutaUra/go-graphql-server/ent"
	"github.com/YutaUra/go-graphql-server/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *ent.Client
}

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client},
	})
}
