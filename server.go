package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"os"
	"testt/db/keystore"
	"testt/graph"
	"testt/graph/generated"
	"testt/handlers"
)

const (
	defaultPort = "5000"
	defaultHost = "http://localhost:" + defaultPort
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	host := os.Getenv("HOST_NAME")
	if host == "" {
		host = defaultHost
	}

	dataStore := keystore.New()

	resolvers := &graph.Resolver{
		Store: dataStore,
		Host:  host,
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))

	httpHandlers := handlers.NewHandler(dataStore)

	http.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	http.HandleFunc("/", httpHandlers.HandleShortUrl)

	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
