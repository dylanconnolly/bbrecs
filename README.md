# Dev tools
For running locally
- Go (https://go.dev/dl/)
- golang migrate CLI (https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- PostgreSQL (https://www.postgresql.org/download/)

# Migrations
- download migrate CLI (https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
```
## To create new migration
migrate create -ext sql -dir postgres/migrations -seq <migration_name>

## To run migration
migrate -database "postgres://localhost:5432/bbrecs?sslmode=disable" -path postgres/migrations up

## To reverse a migration
migrate -database "postgres://localhost:5432/bbrecs?sslmode=disable" -path postgres/migrations down
# can add a version to only run down migrations up to that version
migrate -database "postgres://localhost:5432/bbrecs?sslmode=disable" -path postgres/migrations down 3 (will run all down migrations >= 3)

## To force DB to version
migrate -database "postgres://localhost:5432/bbrecs?sslmode=disable" -path postgres/migrations force <version>

## e.g. `migrate -database "postgres://localhost:5432/bbrecs?sslmode=disable" -path postgres/migrations force 10`
```