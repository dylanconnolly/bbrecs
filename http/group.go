package http

import (
	"net/http"

	"github.com/dylanconnolly/bbrecs/bbrecs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddUserToGroupReqBody struct {
	UserID uuid.UUID `json:"userID"`
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

func (s *Server) handleAddUserToGroup(c *gin.Context, groupID string) {
	var reqBody AddUserToGroupReqBody
	gid, err := uuid.Parse(groupID)
	if err != nil {
		c.String(http.StatusBadRequest, "group ID is invalid uuid: %s", err)
		return
	}

	err = c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.String(http.StatusBadRequest, "request body could not be parsed: %s", err)
		return
	}

	err = s.GroupService.AddUserToGroup(c, gid, reqBody.UserID)
	if err != nil {
		c.String(http.StatusInternalServerError, "error adding user (%s) to group (%s): %s", reqBody.UserID, groupID, err)
		return
	}

	c.String(http.StatusCreated, "successfully added user (%s) to group (%s)", reqBody.UserID, groupID)
}
