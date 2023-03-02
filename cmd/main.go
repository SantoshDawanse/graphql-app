package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	log "github.com/santoshdawanse/graphql-app/config/logger"
	"github.com/santoshdawanse/graphql-app/graph"
	"github.com/santoshdawanse/graphql-app/graph/resolver"
	"github.com/santoshdawanse/graphql-app/server"
	"github.com/sirupsen/logrus"
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

	// starting a server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))
	echo := server.New(srv)
	logrus.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	logrus.Error(echo.Start(":" + port))
}
