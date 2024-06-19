package http

import (
	"net/http"

	"github.com/dylanconnolly/bbrecs/bbrecs"
	"github.com/gin-gonic/gin"
)

type Server struct {
	server       *http.Server
	router       *gin.Engine
	UserService  bbrecs.UserService
	GroupService bbrecs.GroupService
}

func NewServer() *Server {
	s := &Server{
		server: &http.Server{},
		router: NewRouter(),
	}
	// register api routes
	api := s.router.Group("/api")
	s.registerUserApiRoutes(api)
	s.registerGroupApiRoutes(api)

	return s
}

func (s *Server) Run() {
	s.router.Run()
}
