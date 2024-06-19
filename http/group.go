package http

import (
	"net/http"

	"github.com/dylanconnolly/bbrecs/bbrecs"
	"github.com/gin-gonic/gin"
)

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
