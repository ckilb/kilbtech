package route

import (
	"fmt"
	"log"
	"net/http"
)

type Robots struct {
}

func (r *Robots) Path() string {
	return "/robots.txt"
}

func (r *Robots) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		_, err := w.Write([]byte("User-agent: *\nAllow: /"))

		if err != nil {
			log.Println(fmt.Errorf("writing robots.txt response: %w", err))
		}
	})
}

func NewRobots() Route {
	return &Robots{}
}
