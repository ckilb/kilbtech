package route

import (
	"ckilb/kilbtech/tpl"
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
		data := tpl.TemplateData{}

		c.HTML(http.StatusOK, "legal", data)
	}
}

func NewLegal() Route {
	return &Legal{}
}
