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
	generateRoutes(r)

	return r
}

func generateRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "WELCOME -- THIS IS THE ROOT INDEX PAGE")
	})
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.IndentedJSON(http.StatusOK, "Index page of /api")
		})
	}
}
