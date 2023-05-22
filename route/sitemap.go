package route

import (
	"ckilb/kilbtech/blog"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Sitemap struct {
	pages []string
	posts []blog.Post
}

func (r *Sitemap) Path() string {
	return "/sitemap.xml"
}

func (r *Sitemap) Method() string {
	return http.MethodGet
}

func (r *Sitemap) Page() string {
	return ""
}

func (r *Sitemap) Handler() gin.HandlerFunc {
	lines := []string{
		"<?xml version=\"1.0\" encoding=\"UTF-8\"?>",
		"<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">",
	}

	for _, post := range r.posts {
		if !post.IsActive {
			continue
		}

		lines = append(lines, "<url>")
		lines = append(lines, fmt.Sprintf("<loc>https://kilb.tech/%s</loc>", post.Id))
		lines = append(lines, "</url>")
	}

	lines = append(lines, "</urlset>")

	content := strings.Join(lines, "\n")

	return func(c *gin.Context) {
		c.Data(http.StatusOK, "application/xml", []byte(content))
	}
}

func NewSitemap(pages []string, posts []blog.Post) Route {
	return &Sitemap{
		pages: pages,
		posts: posts,
	}
}
