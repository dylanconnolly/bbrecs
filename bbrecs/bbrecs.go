package bbrecs

import (
	"fmt"
	"image/jpeg"
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	InviteCode string    `json:"inviteCode"`
}

type User struct {
	ID          uuid.UUID `json:"id"`
	DisplayName string    `json:"displayName"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	PhoneNumber string    `jsone:"phoneNumber"`
}

type GroupUser struct {
	ID      uuid.UUID `json:"id"`
	UserID  uint
	GroupID uint
}

type Rec struct {
	Title string `json:"title"`
	// Tags enum #allow user creation?
	AuthorComment Comment `json:"authorComment"`
	AuthorPhotos  []Photo `json:"authorPhotos"`
	UserPhotos    []Photo `json:"userPhotos"`
	// AuthorPhotos #grab image from google api
}

type Photo struct {
	Data []byte `json:"data"`
	// DefaultQuality = 75
	Options jpeg.Options
}

type Comment struct {
	User      User      `json:"user"`
	Comment   string    `json:"string"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUser(firstName string, lastName string, phoneNumber string) User {
	// can validate inputs (fname, lname, pnumber) here before sending to DB layer
	user := User{
		ID:          uuid.New(),
		DisplayName: fmt.Sprintf("%s %s", firstName, lastName),
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
	}

	return user
}

func NewGroup(name string) Group {
	group := Group{
		ID:         uuid.New(),
		Name:       name,
		InviteCode: "testCode",
	}

	return group
}
