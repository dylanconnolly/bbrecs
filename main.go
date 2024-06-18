package main

import (
	"fmt"
	"log"

	"github.com/dylanconnolly/bbrecs/http"
	"github.com/dylanconnolly/bbrecs/postgres"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading environment configuration: %s", err)
	}
	fmt.Println("BB RECS")
	s := http.NewServer()
	s.UserService = postgres.NewUserService(postgres.CreatePostgresConnPool())
	s.Run()
}
