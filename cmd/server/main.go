package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"sandbox-gql/graph"
	"sandbox-gql/internal/db"

	_ "github.com/go-sql-driver/mysql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {

	client, err := db.Client()
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	ctx := context.Background()
	u, err := client.Todo.
		Create().
		SetName("John").
		SetEmail("hogehoge@gmail.com").
		SetDone(false).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
