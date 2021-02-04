package main

import (
	"go-graphql-test/config"
	"go-graphql-test/database"
	"go-graphql-test/graph"
	"go-graphql-test/graph/generated"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "8081"

func init() {
	log.Print("Loading .env")
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	config := config.New()

	database.Connect(config.Mongo.Uri)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	playgroundHandler := playground.Handler("GraphQL playground", "/graphql")
	http.Handle("/", playgroundHandler)
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
