package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"sandbox-gql/ent"
	"sandbox-gql/graph"

    _ "github.com/go-sql-driver/mysql"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {

	entOptions := []ent.Option{}
	entOptions = append(entOptions, ent.Debug())
	client, err := ent.Open(dialect.MySQL, "usr:usrpwd@tcp(db:3306)/db", entOptions...)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
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
