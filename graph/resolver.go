package graph

import (
	"sandbox-gql/ent"
	"github.com/go-redis/redis"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ 
	client *ent.Client 
	redis *redis.Client
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client, redis *redis.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{client, redis},
	})
}

