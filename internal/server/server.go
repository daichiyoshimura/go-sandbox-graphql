package server

import (
	"fmt"
	"net/http"
	"sandbox-gql/ent"
	"sandbox-gql/graph"
	"sandbox-gql/internal/env"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-redis/redis"
)

func NewServer(dbClient *ent.Client, redisClient *redis.Client) *handler.Server {
	return handler.NewDefaultServer(graph.NewSchema(dbClient, redisClient))
}

func ListenAndServe(srvvars *env.Server, handler http.Handler) error {
	return http.ListenAndServe(fmt.Sprintf("%s:%s", srvvars.Host(), srvvars.Port()), handler)
}
