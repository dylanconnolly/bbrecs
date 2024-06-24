package bbrecs

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type GroupService interface {
	CreateGroup(c context.Context, name string) (*Group, error)
}

type GroupUserService interface {
	GetGroupUsers(c context.Context, GroupID uuid.UUID) ([]*User, error)
	AddUserToGroup(c context.Context, GroupID uuid.UUID, UserID uuid.UUID) error
	RemoveUserFromGroup(c context.Context, GroupID uuid.UUID, UserID uuid.UUID) error
}

type Group struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
