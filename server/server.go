package server

import (
	"ckilb/kilbtech/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Server struct {
	port   int
	routes []route.Route
}

func (s *Server) Start() error {
	r := gin.Default()

	for _, r := range s.routes {
		http.Handle(r.Path(), r.Handler())
	}

	if err := r.Run("127.0.0.1:" + strconv.Itoa(s.port)); err != nil {
		return fmt.Errorf("starting server: %w", err)
	}

	return nil
}

func NewServer(port int, routes []route.Route) *Server {
	return &Server{
		port:   port,
		routes: routes,
	}
}
