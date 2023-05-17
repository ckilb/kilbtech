package route

import (
	"ckilb/kilbtech/dto"
	"ckilb/kilbtech/tpl"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AboutTemplateData struct {
	AssetCacheId    string
	Projects        []dto.Project
	Headline        string
	SubHeadline     string
	MetaDescription string
}

type Legal struct {
}

func (r *Legal) Path() string {
	return "/legal"
}

func (r *Legal) Method() string {
	return http.MethodGet
}

func (r *Legal) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := tpl.TemplateData{
			MetaDescription: "Legal notice (Impressum) for kilb.tech",
		}

		c.HTML(http.StatusOK, "legal", data)
	}
}

func NewLegal() Route {
	return &Legal{}
}
