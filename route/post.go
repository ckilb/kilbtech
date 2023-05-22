package route

import (
	"ckilb/kilbtech/blog"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
	post blog.Post
}

func (r *Post) Path() string {
	return "/" + r.post.Id
}

func (r *Post) Method() string {
	return http.MethodGet
}

func (r *Post) Page() string {
	return "posts/" + r.post.Id
}

func (r *Post) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, r.Page(), gin.H{
			"Post": r.post,
		})
	}
}

func NewPost(p blog.Post) Route {
	return &Post{p}
}
