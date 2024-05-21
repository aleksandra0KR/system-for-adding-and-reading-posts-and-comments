package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v10"
	"log"
	"net/http"
	"os"
	"system-for-adding-and-reading-posts-and-comments/graph"
	"system-for-adding-and-reading-posts-and-comments/innternal/repository"
)

func main() {
	DB := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "OzonContest",
		Addr:     "localhost:5436",
	})

	defer DB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	repo := repository.NewRepository(DB)
	c := graph.Config{Resolvers: &graph.Resolver{
		Repository: repo}}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
