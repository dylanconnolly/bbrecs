package main

import (
	"log"
	"os"

	"github.com/dylanconnolly/bbrecs/http"
	"github.com/dylanconnolly/bbrecs/postgres"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading environment configuration: %s", err)
	}
	s := http.NewServer()
	postgresURL := os.Getenv("POSTGRES_URL")
	if postgresURL == "" {
		log.Fatalf("Missing postgres connection url in env")
	}
	dbpool := postgres.CreatePostgresConnPool(postgresURL)
	s.UserService = postgres.NewUserService(dbpool)
	s.GroupService = postgres.NewGroupService(dbpool)
	s.Run()
}
