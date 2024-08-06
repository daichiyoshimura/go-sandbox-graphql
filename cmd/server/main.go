package main

import (
	"context"
	"log"

	"sandbox-gql/internal/db"
	"sandbox-gql/internal/env"
	"sandbox-gql/internal/redis"
	"sandbox-gql/internal/server"

	_ "github.com/go-sql-driver/mysql"
)

var ctx = context.Background()

func main() {

	// env
	dbvars, srvvars, redisVars, err := env.Load()
	if err != nil {
		log.Fatalf("failed loading env vars: %v", err)
	}

	// redis client
	redisClient, err := redis.Client(ctx, redisVars)
	if err != nil {
		log.Fatalf("failed connecting to redis: %v", err)
	}

	// database client
	dbClient, err := db.Client(dbvars)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer dbClient.Close()

	// database migration
	if err := dbClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// server
	srv := server.NewServer(dbClient, redisClient)
	server.DefineRoute(srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", srvvars.Port())
	if err := server.ListenAndServe(srvvars, nil); err != nil {
		log.Fatal(err)
	}
}
