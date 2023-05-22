package tpl

import (
	"ckilb/kilbtech/dto"
	"ckilb/kilbtech/route"
	"embed"
	_ "embed"
	"errors"
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin/render"
	"github.com/google/uuid"
	"html"
	"html/template"
	"io"
)

//go:embed **/**/*.tmpl **/*.tmpl *.tmpl
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

func (r *renderer) addPage(name string, paths ...string) error {
	partials, err := partialPaths()

	if err != nil {
		return err
	}

	files := append([]string{"layout.tmpl"}, partials...)
	files = append(files, paths...)

	if err := r.addFromFsFiles("layout.tmpl", name, files); err != nil {
		return fmt.Errorf("add from fs files: %w", err)
	}

	return nil
}

func (r *renderer) addFromFsFiles(base string, name string, files []string) error {
	assetId := uuid.NewString()
	funcMap := template.FuncMap{
		"assetId": func() string {
			return assetId
		},
		"slice": func(args ...interface{}) []interface{} {
			return args
		},
		"escape": func(arg string) string {
			return html.EscapeString(arg)
		},
		"raw": func(arg string) template.HTML {
			return template.HTML(arg)
		},
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}

				dict[key] = values[i+1]
			}
			return dict, nil
		},
	}

	tmpl, err := template.New(base).Funcs(funcMap).ParseFS(filesystem, files...)

	if err != nil {
		return fmt.Errorf("parse file system template: %w", err)
	}

	r.Add(name, tmpl)

	return nil
}

func NewRenderer(routes []route.Route) (render.HTMLRender, error) {
	mr := multitemplate.NewRenderer()
	renderer := &renderer{}
	renderer.Renderer = mr

	for _, r := range routes {
		if r.Page() == "" {
			continue
		}

		if err := renderer.addPage(r.Page(), "pages/"+r.Page()+".tmpl"); err != nil {
			return renderer, fmt.Errorf("add page %s: %w", r.Page(), err)
		}
	}

	return renderer, nil
}
