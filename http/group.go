package http

import (
	"net/http"

	"github.com/dylanconnolly/bbrecs/bbrecs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddUserToGroupReqBody struct {
	GroupID uuid.UUID `json:"groupID"`
	UserID  uuid.UUID `json:"userID"`
}

func (s *Server) handleCreateGroup(c *gin.Context) {
	var group bbrecs.Group

	err := c.ShouldBindJSON(&group)
	if err != nil {
		c.String(http.StatusBadRequest, "request body could not be parsed into Group struct %s", err)
	}

	createdGroup, err := s.GroupService.CreateGroup(c, group.Name)

	if err != nil {
		c.String(http.StatusInternalServerError, "error creating Group in database %s", err)
	}

	c.IndentedJSON(http.StatusCreated, createdGroup)
}

func (s *Server) handleAddUserToGroup(c *gin.Context) {
	var reqBody AddUserToGroupReqBody

	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.String(http.StatusBadRequest, "request body could not be parsed: %s", err)
	}

	err = s.GroupService.AddUserToGroup(c, reqBody.GroupID, reqBody.UserID)
	if err != nil {
		c.String(http.StatusInternalServerError, "error adding user (%s) to group (%s): %s", reqBody.UserID, reqBody.GroupID, err)
	}

	c.String(http.StatusCreated, "successfully added user (%s) to group (%s)", reqBody.UserID, reqBody.GroupID)
}
