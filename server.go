package main

import (
  "log"
  "net/http"
  "os"
  "digitalocean/graph"
  "digitalocean/graph/generated"

  "github.com/99designs/gqlgen/graphql/handler"
  "github.com/99designs/gqlgen/graphql/playground"
  "github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
  err := godotenv.Load(); if err != nil {
    log.Fatal("Error loading .env file")
  }

  port := os.Getenv("PORT")
  if port == "" {
     port = defaultPort
  }

    Database := graph.Connect()
    srv := handler.NewDefaultServer(
            generated.NewExecutableSchema(
                    generated.Config{
                        Resolvers: &graph.Resolver{
                            DB: Database,
                        },
                    }),
        )

  http.Handle("/", playground.Handler("GraphQL playground", "/query"))
  http.Handle("/query", srv)

  log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
  log.Fatal(http.ListenAndServe(":"+port, nil))
}

