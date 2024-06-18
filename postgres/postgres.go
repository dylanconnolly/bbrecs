package postgres

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePostgresConnPool() *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("Unable to create postgres connection pool: %s", err)
	}
	log.Printf("connected to postgres: %s", os.Getenv("POSTGRES_URL"))
	return dbpool
}
