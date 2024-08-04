package main

import (
	"context"
	"log"

	"os"

	"sandbox-gql/internal/db"
	"sandbox-gql/internal/server"

	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

func main() {

	// database client
	dbClient, err := db.Client()
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer dbClient.Close()

	// database migration
	if err := dbClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// env
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// server
	srv := server.NewServer(dbClient)
	server.DefineRoute(srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	if err := server.ListenAndServe(":", port, nil); err != nil {
		log.Fatal(err)
	}
}
