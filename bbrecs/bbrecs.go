package bbrecs

import (
	"context"
	"image/jpeg"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(c context.Context, user *User) (*User, error)
	// GetUserByID(userID uuid.UUID) (*User, error)
	// GetGroupUsers(groupID uuid.UUID) ([]*User, error)
}

func GenerateUser(userData NewUserFields) (*User, error) {
	user := User{
		NewUserFields: userData,
	}

	return &user, nil
}

type Group struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	InviteCode string    `json:"inviteCode"`
}

type User struct {
	ID uuid.UUID `json:"id"`
	NewUserFields
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type NewUserFields struct {
	FirstName   string `json:"firstName" db:"first_name"`
	LastName    string `json:"lastName" db:"last_name"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
}
type BaseRec struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	// Tags enum #allow user creation?
	Author      User    `json:"author"`
	Description string  `json:"description"`
	UserPhotos  []Photo `json:"userPhotos"`
	// AuthorPhotos #grab image from google api
}

type Photo struct {
	ID   uuid.UUID `json:"id"`
	Data []byte    `json:"data"`
	// DefaultQuality = 75
	Options jpeg.Options
}

type Comment struct {
	ID        uuid.UUID `json:"id"`
	User      User      `json:"user"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
