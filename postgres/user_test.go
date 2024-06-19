package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/dylanconnolly/bbrecs/bbrecs"
)

func TestCreateUser(t *testing.T) {
	us := NewUserService(CreatePostgresConnPool("postgres://localhost:5432/bbrecs?sslmode=disable"))

	user := &bbrecs.User{
		NewUserFields: bbrecs.NewUserFields{
			FirstName:   "test",
			LastName:    "user",
			PhoneNumber: "1112223333",
		},
	}

	createdUser, err := us.CreateUser(context.Background(), user)

	if err != nil {
		t.Errorf("err is not nil: %v", err)
	}

	fmt.Printf("new user: %+v", createdUser)

}

func TestGetUsers(t *testing.T) {
	us := NewUserService(CreatePostgresConnPool("postgres://localhost:5432/bbrecs?sslmode=disable"))

	users, err := us.GetUsers(context.Background())

	if err != nil {
		t.Errorf("err is not nil: %v", err)
	}

	if len(users) == 0 {
		t.Errorf("length is 0, expected to have at least 1 entry")
	}
}
