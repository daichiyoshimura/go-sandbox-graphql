package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func DefineRoute(server *handler.Server) {
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", server)
}