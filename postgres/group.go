package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/dylanconnolly/bbrecs/bbrecs"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupService struct {
	db *pgxpool.Pool
}

type GroupUserService struct {
	db *pgxpool.Pool
}

func NewGroupService(dbpool *pgxpool.Pool) *GroupService {
	return &GroupService{db: dbpool}
}

func NewGroupUserService(dbpool *pgxpool.Pool) *GroupUserService {
	return &GroupUserService{db: dbpool}
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

func (s *GroupUserService) GetGroupUsers(c context.Context, GroupID uuid.UUID) ([]*bbrecs.User, error) {
	tx, err := s.db.Begin(c)

	if err != nil {
		return nil, err
	}
	defer tx.Rollback(c)

	query := `
		SELECT users.* FROM group_users
		JOIN users ON group_users.user_id = users.id
		WHERE group_users.group_id = $1;
	`

	rows, err := tx.Query(c, query, GroupID)
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

func (s *GroupUserService) AddUserToGroup(c context.Context, GroupID uuid.UUID, UserID uuid.UUID) error {
	tx, err := s.db.Begin(c)

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
		if pgErr, ok := err.(*pgconn.PgError); ok {
			// suppress error if group to user relationship already exists
			if pgErr.Code == "23505" {
				return nil
			}
		}
		log.Printf("error adding user to group - %s", err)
		return err
	}

	err = tx.Commit(c)
	if err != nil {
		log.Printf("error committing transaction in AddUserToGroup - %s", err)
		return err
	}

	return nil
}

func (s *GroupUserService) RemoveUserFromGroup(c context.Context, GroupID uuid.UUID, UserID uuid.UUID) error {
	tx, err := s.db.Begin(c)

	if err != nil {
		return err
	}
	defer tx.Rollback(c)

	query := `
		DELETE FROM group_users
		WHERE group_id=$1 AND user_id=$2
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
