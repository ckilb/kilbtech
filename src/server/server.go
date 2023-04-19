package server

import (
	"ckilb/kilbtech/route"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	port   int
	routes []route.Route
}

func (s *Server) Start() error {
	for _, r := range s.routes {
		http.Handle(r.Path(), r.Handler())
	}

	if err := http.ListenAndServe("127.0.0.1:"+strconv.Itoa(s.port), nil); err != nil {
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
