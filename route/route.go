package route

import (
	"github.com/gin-gonic/gin"
)

type Route interface {
	Path() string
	Handler() gin.HandlerFunc
	Method() string
	Page() string
	Templates() []string
}
