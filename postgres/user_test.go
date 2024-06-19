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
		t.Errorf("err is not nil, %v", err)
	}

	fmt.Printf("new user: %+v", createdUser)

}
