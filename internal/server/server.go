package server

import (
	"fmt"
	"net/http"
	"sandbox-gql/ent"
	"sandbox-gql/graph"
	"sandbox-gql/internal/env"

	"github.com/99designs/gqlgen/graphql/handler"
)

func NewServer(dbClient *ent.Client) *handler.Server {
	return handler.NewDefaultServer(graph.NewSchema(dbClient))
}

func ListenAndServe(srvvars *env.Server, handler http.Handler) error {
	return http.ListenAndServe(fmt.Sprintf("%s:%s", srvvars.Host(), srvvars.Port()), handler)
}
