package http

import (
	"net/http"

	"github.com/dylanconnolly/bbrecs/bbrecs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) handleGetUserByID(c *gin.Context, userID string) {
	var user bbrecs.User

	user, err := s.UserService.GetUserByID(c, uid)

	if err != nil {
		c.String(http.StatusInternalServerError, "couldn't get user: %s", err)
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (s *Server) handleGetUsers(c *gin.Context) {
	var users []*bbrecs.User

	users, err := s.UserService.GetUsers(c)

	if err != nil {
		c.String(http.StatusInternalServerError, "oops couldnt get users: %s", err)
	}

	c.IndentedJSON(http.StatusOK, users)
}

func (s *Server) handleUpdateUser(c *gin.Context, userID string) {
	var fields bbrecs.UserUpdateFields

	uid, err := uuid.Parse(userID)
	if err != nil {
		c.String(http.StatusBadRequest, "could not parse user ID (must be UUID): %s", err)
		return
	}

	err = c.ShouldBindJSON(&fields)
	if err != nil {
		c.String(http.StatusBadRequest, "request body could not be parsed into User struct %s", err)
	}

	user, err := s.UserService.UpdateUser(c, uid, fields)

	if err != nil {
		c.String(http.StatusInternalServerError, "oops couldnt get users: %s", err)
	}

	c.IndentedJSON(http.StatusOK, user)
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

func (s *Server) handleGetUserGroups(c *gin.Context, userID string) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		c.String(http.StatusBadRequest, "could not parse user ID (must be UUID): %s", err)
		return
	}

	groups, err := s.UserService.GetUserGroups(c, uid)
	if err != nil {
		c.String(http.StatusInternalServerError, "error getting groups: %s", err)
	}

	c.IndentedJSON(http.StatusOK, groups)
}
