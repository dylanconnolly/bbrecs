package postgres

import (
	"context"
	"fmt"

	"github.com/dylanconnolly/bbrecs/bbrecs"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

	err = tx.Commit(c)
	if err != nil {
		return nil, fmt.Errorf("error committing tx: %s", err)
	}

	return &newUser, nil
}

// TODO: pass struct with nil values so they don't get updated? Or always expect to override all columns with new values even if they are not changing
func (us *UserService) UpdateUser(c context.Context, userID uuid.UUID, fields bbrecs.UserUpdateFields) (*bbrecs.User, error) {
	tx, err := us.db.Begin(c)

	if err != nil {
		return nil, err
	}
	defer tx.Rollback(c)

	query := `
		UPDATE users
		SET first_name = $2,
			last_name = $3,
			phone_number = $4
		WHERE id = $1
		RETURNING id, first_name, last_name, phone_number, created_at, updated_at;
	`

	// _, err = tx.Exec(c, query, user.FirstName, user.LastName, user.PhoneNumber)
	var user bbrecs.User
	err = tx.QueryRow(c, query, userID, fields.FirstName, fields.LastName, fields.PhoneNumber).Scan(&user.ID, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		fmt.Printf("%+v", user)
		return nil, fmt.Errorf("error with query: %v", err)
	}

	err = tx.Commit(c)
	if err != nil {
		return nil, fmt.Errorf("error committing tx: %s", err)
	}

	return &user, nil
}

func (us *UserService) GetUsers(c context.Context) ([]*bbrecs.User, error) {
	tx, err := us.db.Begin(c)

	if err != nil {
		return nil, err
	}
	defer tx.Rollback(c)

	query := `SELECT * FROM users;`

	rows, err := tx.Query(c, query)
	if err != nil {
		return nil, err
	}
	users, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*bbrecs.User, error) {
		user, err := pgx.RowToStructByName[bbrecs.User](row)
		if err != nil {
			return nil, err
		}
		return &user, err
	})

	tx.Commit(c)

	return users, err
}

func (us *UserService) GetUserGroups(c context.Context, userID uuid.UUID) ([]bbrecs.Group, error) {
	tx, err := us.db.Begin(c)

	if err != nil {
		return nil, err
	}
	defer tx.Rollback(c)

	query := `
		SELECT groups.id, groups.name, groups.created_at, groups.updated_at FROM groups 
		JOIN group_users ON groups.id = group_users.group_id
		WHERE group_users.user_id=$1;
	`

	rows, err := tx.Query(c, query, userID)
	if err != nil {
		return nil, err
	}
	groups, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (bbrecs.Group, error) {
		var group bbrecs.Group
		err := row.Scan(&group.ID, &group.Name, &group.CreatedAt, &group.UpdatedAt)
		return group, err
	})

	return groups, err
}
