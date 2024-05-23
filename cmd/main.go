package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"system-for-adding-and-reading-posts-and-comments/graph"
	"system-for-adding-and-reading-posts-and-comments/innternal/repository/inMemory"
	postgres2 "system-for-adding-and-reading-posts-and-comments/innternal/repository/postgres"
)

const (
	host     = "ozon-contest"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "OzonContest"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
}
func main() {
	typeStorage, exists := os.LookupEnv("STORAGE")
	if !exists {
		log.Fatal("STORAGE environment variable not set")
	}
	var c graph.Config
	if typeStorage == "in-memory" {
		repo := inMemory.NewInMemoryRepository()
		c = graph.Config{Resolvers: &graph.Resolver{Repository: repo}}
	} else {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatalf("connection failed: %s", err.Error())
		}
		defer db.Close()
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			log.Fatal(err)
		}
		migration, err := migrate.NewWithDatabaseInstance("file://database/migration",
			dbname, driver)
		if err != nil {
			log.Fatal(err)
		}
		if err := migration.Up(); err != nil {
			if !errors.Is(err, migrate.ErrNoChange) {
				log.Fatal(err)
			}
		}

		repo := postgres2.NewRepository(db)
		c = graph.Config{Resolvers: &graph.Resolver{Repository: repo}}
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":"+"8080", nil))
}
