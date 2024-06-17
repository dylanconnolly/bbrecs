package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *http.Server
	Router *gin.Engine
}

func NewServer() *Server {
	s := &Server{
		server: &http.Server{},
		Router: NewRouter(),
	}

	return s
}

func (s *Server) Serve() {
	s.Router.Run()
}
