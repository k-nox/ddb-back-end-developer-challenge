package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/k-nox/ddb-backend-developer-challenge/app"
	"github.com/k-nox/ddb-backend-developer-challenge/graph"
	"github.com/k-nox/ddb-backend-developer-challenge/graph/generated"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	// setup the app & db
	app, err := app.New("db/data.db", "db/migrations")
	if err != nil {
		log.Fatal(err)
	}
	defer app.CloseDB()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.New(app)}))
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL Playground", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
