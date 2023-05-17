package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Robots struct {
}

func (r *Robots) Path() string {
	return "/robots.txt"
}

func (r *Robots) Method() string {
	return http.MethodGet
}
func (r *Robots) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "User-agent: *\nAllow: /")
	}
}

func NewRobots() Route {
	return &Robots{}
}
