package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	log "github.com/santoshdawanse/graphql-app/config/logger"
	"github.com/santoshdawanse/graphql-app/graph"
	"github.com/santoshdawanse/graphql-app/graph/resolver"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

const defaultPort = "8080"

func init() {
	log.InitializeLogger()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logrus.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	logrus.Error(http.ListenAndServe(":"+port, nil))
}
