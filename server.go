package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	"github.com/Swejal08/go-ggqlen/directives"
	graph "github.com/Swejal08/go-ggqlen/graph/resolvers"
	resolvers "github.com/Swejal08/go-ggqlen/graph/resolvers"
	"github.com/Swejal08/go-ggqlen/initializer"
	"github.com/Swejal08/go-ggqlen/middleware"
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

	router := chi.NewRouter()

	router.Use(middleware.UserMiddleware())

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}, Directives: resolvers.DirectiveRoot{CheckUserIdExists: directives.CheckUserId()}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
