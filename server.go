package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/YutaUra/go-graphql-server/ent"
	"github.com/YutaUra/go-graphql-server/graph"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	client, err := ent.Open(dialect.Postgres, fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("POSTGRES_HOSTNAME"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PASSWORD")))
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	srv := handler.NewDefaultServer(graph.NewSchema(client))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
