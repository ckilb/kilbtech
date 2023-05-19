package route

import (
	"ckilb/kilbtech/blog"
	"ckilb/kilbtech/tpl"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostTemplateData struct {
	tpl.TemplateData
	Post *blog.Post
}

type Post struct {
	post *blog.Post
}

func (r *Post) Path() string {
	return "/" + r.post.Id
}

func (r *Post) Method() string {
	return http.MethodGet
}

func (r *Post) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := PostTemplateData{
			Post: r.post,
		}

		c.HTML(http.StatusOK, "posts/"+r.post.Id, data)
	}
}

func NewPost(p *blog.Post) Route {
	return &Post{p}
}
