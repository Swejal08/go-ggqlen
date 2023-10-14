package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	graph "github.com/Swejal08/go-ggqlen/graph/resolvers"
	resolvers "github.com/Swejal08/go-ggqlen/graph/resolvers"
	"github.com/Swejal08/go-ggqlen/initializer"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func init() {
	initializer.LoadEnvVariables()
	initializer.InitializeDatabase()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
