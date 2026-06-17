package route

import (
	"ckilb/kilbtech/blog"
	"fmt"
	"html"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Rss struct {
	posts []blog.Post
}

func (r *Rss) Path() string {
	return "/rss.xml"
}

func (r *Rss) Method() string {
	return http.MethodGet
}

func (r *Rss) Page() string {
	return ""
}

func (r *Rss) Templates() []string {
	return []string{}
}

func (r *Rss) Handler() gin.HandlerFunc {
	active := []blog.Post{}
	for _, p := range r.posts {
		if p.IsActive {
			active = append(active, p)
		}
	}

	sort.Slice(active, func(i, j int) bool {
		return active[i].Date > active[j].Date
	})

	pubDate := func(dateStr string) string {
		t, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return ""
		}
		return t.UTC().Format(time.RFC1123Z)
	}

	lastBuildDate := ""
	if len(active) > 0 {
		lastBuildDate = pubDate(active[0].Date)
	}

	lines := []string{
		`<?xml version="1.0" encoding="UTF-8"?>`,
		`<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">`,
		`<channel>`,
		`<title>Christian Kilb — Blog</title>`,
		`<link>https://kilb.tech/blog</link>`,
		`<description>Software engineering thoughts by Christian Kilb</description>`,
		`<language>en</language>`,
		fmt.Sprintf(`<lastBuildDate>%s</lastBuildDate>`, lastBuildDate),
		`<atom:link href="https://kilb.tech/rss.xml" rel="self" type="application/rss+xml"/>`,
	}

	for _, p := range active {
		lines = append(lines,
			`<item>`,
			fmt.Sprintf(`<title>%s</title>`, html.EscapeString(p.Title)),
			fmt.Sprintf(`<link>https://kilb.tech/%s</link>`, p.Id),
			fmt.Sprintf(`<guid isPermaLink="true">https://kilb.tech/%s</guid>`, p.Id),
			fmt.Sprintf(`<description>%s</description>`, html.EscapeString(p.Subtitle)),
			fmt.Sprintf(`<pubDate>%s</pubDate>`, pubDate(p.Date)),
			`</item>`,
		)
	}

	lines = append(lines,
		`</channel>`,
		`</rss>`,
	)

	content := strings.Join(lines, "\n")

	return func(c *gin.Context) {
		c.Data(http.StatusOK, "application/rss+xml; charset=utf-8", []byte(content))
	}
}

func NewRss(posts []blog.Post) Route {
	return &Rss{posts: posts}
}
