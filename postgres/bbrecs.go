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

func (us *UserService) GetUsers(c context.Context) ([]bbrecs.User, error) {
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
	users, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (bbrecs.User, error) {
		var user bbrecs.User
		err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)
		return user, err
	})

	return users, err
}

type GroupService struct {
	db *pgxpool.Pool
}

func NewGroupService(dbpool *pgxpool.Pool) *GroupService {
	return &GroupService{db: dbpool}
}

func (gs *GroupService) CreateGroup(c context.Context, name string) (*bbrecs.Group, error) {
	tx, err := gs.db.Begin(c)

	if err != nil {
		return nil, err
	}
	defer tx.Rollback(c)

	query := `
		INSERT INTO groups (name)
		VALUES ($1)
		RETURNING id, name, created_at, updated_at;
	`

	// _, err = tx.Exec(c, query, user.FirstName, user.LastName, user.PhoneNumber)
	var newGroup bbrecs.Group
	err = tx.QueryRow(c, query, name).Scan(&newGroup.ID, &newGroup.Name, &newGroup.CreatedAt, &newGroup.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error with query: %v", err)
	}

	err = tx.Commit(c)
	if err != nil {
		return nil, fmt.Errorf("error committing tx: %s", err)
	}

	return &newGroup, nil
}

func (gs *GroupService) AddUserToGroup(c context.Context, GroupID uuid.UUID, UserID uuid.UUID) error {
	tx, err := gs.db.Begin(c)

	if err != nil {
		return err
	}
	defer tx.Rollback(c)

	query := `
		INSERT INTO group_users (group_id, user_id)
		VALUES ($1, $2);
	`

	_, err = tx.Exec(c, query, GroupID, UserID)

	if err != nil {
		return err
	}

	err = tx.Commit(c)
	if err != nil {
		return fmt.Errorf("error committing tx: %s", err)
	}

	return nil
}
