package tpl

import (
	"ckilb/kilbtech/dto"
	"embed"
	_ "embed"
	"fmt"
	"github.com/google/uuid"
	"io"
	"text/template"
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
}

func (r *renderer) Render(wr io.Writer, page string, data TemplateData) error {
	partials, err := r.partialPaths()

	if err != nil {
		return fmt.Errorf("get partial paths: %w", err)
	}

	data.AssetId = uuid.NewString()

	files := append([]string{"layout.html"}, partials...)
	files = append(files, "pages/"+page+".html")

	tmpl, err := template.ParseFS(filesystem, files...)

	if err != nil {
		return fmt.Errorf("parsing templates: %w", err)
	}

	if err = tmpl.Execute(wr, data); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	return nil
}

func (r *renderer) partialPaths() ([]string, error) {
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

func NewRenderer() Renderer {
	return &renderer{}
}
