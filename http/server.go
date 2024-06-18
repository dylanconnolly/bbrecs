package http

import (
	"net/http"

	"github.com/dylanconnolly/bbrecs/bbrecs"
	"github.com/gin-gonic/gin"
)

type Server struct {
	server      *http.Server
	Router      *gin.Engine
	UserService bbrecs.UserService
}

func NewServer() *Server {
	s := &Server{
		server: &http.Server{},
		Router: NewRouter(),
	}

	return s
}

func (s *Server) Serve() {
	s.run()
}
