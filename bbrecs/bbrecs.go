package bbrecs

import (
	"image/jpeg"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user *User) (*User, error)
	GetUserByID(userID uuid.UUID) (*User, error)
	GetGroupUsers(groupID uuid.UUID) ([]*User, error)
}

func GenerateUser(userData NewUserFields) (*User, error) {
	user := User{
		ID:            uuid.New(),
		NewUserFields: userData,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
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
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type NewUserFields struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
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
