package postgres

import (
	"context"
	"fmt"

	"github.com/dylanconnolly/bbrecs/bbrecs"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	db *pgxpool.Pool
}

func NewUserService(dbpool *pgxpool.Pool) *UserService {
	return &UserService{db: dbpool}
}

func (us *UserService) CreateUser(c context.Context, user *bbrecs.User) (*bbrecs.User, error) {
	tx, err := us.db.Begin(c)

	if err != nil {
		return nil, err
	}
	defer tx.Rollback(c)

	query := `
		INSERT INTO users (first_name,last_name,phone_number)
		VALUES ($1, $2, $3)
		RETURNING id, first_name, last_name, phone_number, created_at, updated_at;
	`

	// _, err = tx.Exec(c, query, user.FirstName, user.LastName, user.PhoneNumber)
	var newUser bbrecs.User
	err = tx.QueryRow(c, query, user.FirstName, user.LastName, user.PhoneNumber).Scan(&newUser.ID, &newUser.FirstName, &newUser.LastName, &newUser.PhoneNumber, &newUser.CreatedAt, &newUser.UpdatedAt)
	if err != nil {
		fmt.Printf("%+v", user)
		return nil, fmt.Errorf("error with query: %v", err)
	}

	if err != nil {
		return nil, fmt.Errorf("error scanning into struct: %s", err)
	}

	err = tx.Commit(c)
	if err != nil {
		return nil, fmt.Errorf("error committing tx: %s", err)
	}

	return &newUser, nil
}
