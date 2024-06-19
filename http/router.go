package http

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(cors.Default())
	r.Use(gin.Recovery())

	return r
}

func (s *Server) registerUserApiRoutes(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Index page of /api")
	})

	r.GET("/users", func(c *gin.Context) {
		s.handleGetUsers(c)
	})

	r.POST("/users", func(c *gin.Context) {
		s.handleCreateUser(c)
	})
}

func (s *Server) registerGroupApiRoutes(r *gin.RouterGroup) {
	r.POST("/groups", func(c *gin.Context) {
		s.handleCreateGroup(c)
	})
}
