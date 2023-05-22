package route

import (
	"ckilb/kilbtech/blog"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Blog struct {
	posts []blog.Post
}

func (r *Blog) Path() string {
	return "/blog"
}

func (r *Blog) Method() string {
	return http.MethodGet
}

func (r *Blog) Page() string {
	return "blog"
}

func (r *Blog) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog", gin.H{
			"posts": r.posts,
		})
	}
}

func NewBlog(posts []blog.Post) Route {
	return &Blog{posts: posts}
}
