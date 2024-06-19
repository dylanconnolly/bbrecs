package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePostgresConnPool(url string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to create postgres connection pool to '%s': %v", url, err)
	}
	log.Printf("connected to postgres: '%s'", url)
	return dbpool
}
