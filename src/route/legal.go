package route

import (
	"ckilb/kilbtech/dto"
	"ckilb/kilbtech/tpl"
	"fmt"
	"log"
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
	renderer tpl.Renderer
}

func (r *Legal) Path() string {
	return "/legal"
}

func (r *Legal) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		data := tpl.TemplateData{
			MetaDescription: "Legal notice (Impressum) for kilb.tech",
		}

		if err := r.renderer.Render(w, "legal", data); err != nil {
			log.Println(fmt.Errorf("executing legal template: %w", err))
		}
	})
}

func NewLegal(renderer tpl.Renderer) Route {
	return &Legal{renderer: renderer}
}
