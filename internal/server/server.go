package server

import (
	"fmt"
	"net/http"
	"sandbox-gql/ent"
	"sandbox-gql/graph"

	"github.com/99designs/gqlgen/graphql/handler"
)

func NewServer(dbClient *ent.Client) *handler.Server {
	return handler.NewDefaultServer(graph.NewSchema(dbClient))
}

func ListenAndServe(server, port string, handler http.Handler) error {
	return http.ListenAndServe(fmt.Sprintf("%s:%s", server, port), handler)
}
