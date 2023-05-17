package tpl

import (
	"ckilb/kilbtech/dto"
	"embed"
	_ "embed"
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin/render"
	"html/template"
	"io"
)

//go:embed **/*.html *.html
var filesystem embed.FS

type TemplateData struct {
	AssetId         string
	Headline        string
	SubHeadline     string
	MetaDescription string
	Projects        []dto.Project
	TeaserImage     string
}

type Renderer interface {
	Render(wr io.Writer, page string, data TemplateData) error
}
type renderer struct {
	multitemplate.Renderer
}

func partialPaths() ([]string, error) {
	var files []string
	dir, err := filesystem.ReadDir("partials")

	if err != nil {
		return []string{}, fmt.Errorf("read partial dir: %w", err)
	}

	for _, file := range dir {
		files = append(files, "partials/"+file.Name())
	}

	return files, nil

}

func (r *renderer) addPage(name string) error {
	partials, err := partialPaths()

	if err != nil {
		return err
	}

	files := append([]string{"layout.html"}, partials...)
	files = append(files, "pages/"+name+".html")

	tmpl, err := template.ParseFS(filesystem, files...)

	if err != nil {
		return fmt.Errorf("parse file system template: %v", err)
	}

	r.Add(name, tmpl)

	return nil
}

func NewRenderer() (render.HTMLRender, error) {
	mr := multitemplate.NewRenderer()
	r := &renderer{}
	r.Renderer = mr

	pages := []string{"home", "spryker", "legal"}

	for _, page := range pages {
		if err := r.addPage("home"); err != nil {
			return r, fmt.Errorf("add page %s: %v", page, err)
		}
	}

	return r, nil
}
