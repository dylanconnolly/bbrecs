package http

import (
	"net/http"

	"github.com/dylanconnolly/bbrecs/bbrecs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(cors.Default())
	r.Use(gin.Recovery())

	return r
}

func (s *Server) run() {
	s.generateRoutes(s.Router)
	s.Router.Run()
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

func (s *Server) generateRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "WELCOME -- THIS IS THE ROOT INDEX PAGE")
	})
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.IndentedJSON(http.StatusOK, "Index page of /api")
		})

		api.POST("/users", func(c *gin.Context) {
			s.handleCreateUser(c)
		})
	}
}
