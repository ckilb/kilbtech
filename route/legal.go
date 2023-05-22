package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Legal struct {
}

func (r *Legal) Path() string {
	return "/legal"
}

func (r *Legal) Method() string {
	return http.MethodGet
}

func (r *Legal) Page() string {
	return "legal"
}

func (r *Legal) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "legal", gin.H{})
	}
}

func NewLegal() Route {
	return &Legal{}
}
