package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	server *http.Server
	Router *gin.Engine
}

func NewServer() *Server {
	s := &Server{
		server: &http.Server{},
		Router: gin.New(),
	}

	s.Router.Use(cors.Default())
	s.Router.Use(gin.Recovery())
	s.GenerateRoutes(s.Router)

	return s
}

func (s *Server) GenerateRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "WELCOME -- THIS IS THE ROOT INDEX PAGE")
	})
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/", func(c *gin.Context) {
				c.IndentedJSON(http.StatusOK, "This is the index page of api/v1,")
			})
		}
	}
}

func (s *Server) Serve() {
	s.Router.Run()
}
