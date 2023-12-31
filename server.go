package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Swejal08/go-ggqlen/db"
	"github.com/Swejal08/go-ggqlen/graph"
	resolver "github.com/Swejal08/go-ggqlen/graph/resolvers"
	initializers "github.com/Swejal08/go-ggqlen/initializer"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func init() {
	initializers.LoadEnvVariables()
	db.InitializeDatabase()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
