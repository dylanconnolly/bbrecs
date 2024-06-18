package http

import (
	"net/http"

	"github.com/dylanconnolly/bbrecs/bbrecs"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleGetUsers(c *gin.Context) {
	var users []bbrecs.User

	users, err := s.UserService.GetUsers(c)

	if err != nil {
		c.String(http.StatusInternalServerError, "oops couldnt get users: %s", err)
	}

	c.IndentedJSON(http.StatusOK, users)

}

func (s *Server) handleCreateUser(c *gin.Context) {
	var userData bbrecs.NewUserFields

	err := c.ShouldBindJSON(&userData)
	if err != nil {
		c.String(http.StatusBadRequest, "request body could not be parsed into User struct %s", err)
	}

	user, err := bbrecs.GenerateUser(userData)
	if err != nil {
		c.String(http.StatusInternalServerError, "could not generate user %s", err)
	}

	user, err = s.UserService.CreateUser(c, user)

	if err != nil {
		c.String(http.StatusInternalServerError, "error creating user in database %s", err)
	}

	c.IndentedJSON(http.StatusCreated, user)
}
