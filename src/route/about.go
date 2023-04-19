package route

import (
	"ckilb/kilbtech/dto"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"text/template"
)

type AboutTemplateData struct {
	AssetCacheId string
	Projects     []dto.Project
	Headline     string
	SubHeadline  string
}

type Legal struct {
	tpl *template.Template
}

func (r *Legal) Path() string {
	return "/legal"
}

func (r *Legal) Handler() http.Handler {
	assetCacheId := uuid.NewString()

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		err := r.tpl.Execute(w, &AboutTemplateData{
			AssetCacheId: assetCacheId,
		})

		if err != nil {
			log.Println(fmt.Errorf("executing legal template: %w", err))
		}
	})
}

func NewLegal(tpl *template.Template) Route {
	return &Legal{tpl: tpl}
}
