package main

import (
	"graphql-sample/generated"
	"graphql-sample/model"
	"graphql-sample/resolver"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	model.InitDatabase()
	port := "8080"
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{Resolvers: &resolver.Resolver{}}))
	http.Handle("/", playground.Handler("Graphql playground", "/api/graphql"))
	http.Handle("/api/graphql", srv)
	log.Printf("http://localhost:%s for Graphql platground \n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
