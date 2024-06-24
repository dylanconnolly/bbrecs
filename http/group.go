package http

import (
	"net/http"

	"github.com/dylanconnolly/bbrecs/bbrecs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GroupUserReqBody struct {
	UserID uuid.UUID `json:"userID"`
}

func (s *Server) handleCreateGroup(c *gin.Context) {
	var group bbrecs.Group

	err := c.ShouldBindJSON(&group)
	if err != nil {
		c.String(http.StatusBadRequest, "request body could not be parsed into Group struct %s", err)
		return
	}

	createdGroup, err := s.GroupService.CreateGroup(c, group.Name)

	if err != nil {
		c.String(http.StatusInternalServerError, "error creating Group in database %s", err)
		return
	}

	c.IndentedJSON(http.StatusCreated, createdGroup)
}

func (s *Server) handleGetGroupUsers(c *gin.Context, groupID string) {
	gid, err := uuid.Parse(groupID)
	if err != nil {
		c.String(http.StatusBadRequest, "group ID is invalid (must be UUID): %s", err)
		return
	}

	users, err := s.GroupUserService.GetGroupUsers(c, gid)
	if err != nil {
		c.String(http.StatusInternalServerError, "error getting group users: %s", err)
	}

	c.IndentedJSON(http.StatusOK, users)
}

func (s *Server) handleAddUserToGroup(c *gin.Context, groupID string) {
	var reqBody GroupUserReqBody
	gid, err := uuid.Parse(groupID)
	if err != nil {
		c.String(http.StatusBadRequest, "group ID is invalid (must be UUID): %s", err)
		return
	}

	err = c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.String(http.StatusBadRequest, "error parsing request body: %s", err)
		return
	}

	err = s.GroupUserService.AddUserToGroup(c, gid, reqBody.UserID)
	if err != nil {
		c.String(http.StatusInternalServerError, "an error unknown error occurred adding user (%s) to group (%s)", reqBody.UserID, groupID)
		return
	}

	c.String(http.StatusOK, "successfully added user (%s) to group (%s)", reqBody.UserID, groupID)
}

func (s *Server) handleRemoveUserFromGroup(c *gin.Context, groupID string) {
	var reqBody GroupUserReqBody
	gid, err := uuid.Parse(groupID)
	if err != nil {
		c.String(http.StatusBadRequest, "group ID is invalid (must be UUID): %s", err)
		return
	}

	err = c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.String(http.StatusBadRequest, "request body could not be parsed: %s", err)
		return
	}

	err = s.GroupUserService.RemoveUserFromGroup(c, gid, reqBody.UserID)
	if err != nil {
		c.String(http.StatusInternalServerError, "error removing user (%s) from group (%s): %s", reqBody.UserID, groupID, err)
		return
	}

	c.String(http.StatusOK, "successfully removed user (%s) from group (%s)", reqBody.UserID, groupID)
}
