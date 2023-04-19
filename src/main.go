package main

import (
	"ckilb/kilbtech/route"
	"ckilb/kilbtech/server"
	_ "embed"
	"flag"
	"fmt"
	"text/template"
)

var (
	//go:embed template/layout.html
	layoutTemplate string

	//go:embed template/teaser.html
	teaserTemplate string

	//go:embed template/projects.html
	projectsTemplate string

	//go:embed template/home.html
	homeTemplate string

	//go:embed template/spryker.html
	sprykerTemplate string

	//go:embed template/legal.html
	legalTemplate string
)

func main() {
	port := flag.Int("port", 9090, "port")
	staticPath := flag.String("staticPath", "../static", "staticPath")

	flag.Parse()

	tplHome := template.New("home")
	tplSpryker := template.New("spryker")
	tplLegal := template.New("legal")

	if _, err := tplHome.Parse(layoutTemplate + teaserTemplate + projectsTemplate + homeTemplate); err != nil {
		panic(fmt.Errorf("parsing home template: %w", err))
	}

	if _, err := tplSpryker.Parse(layoutTemplate + teaserTemplate + projectsTemplate + sprykerTemplate); err != nil {
		panic(fmt.Errorf("parsing spryker template: %w", err))
	}

	if _, err := tplLegal.Parse(layoutTemplate + teaserTemplate + projectsTemplate + legalTemplate); err != nil {
		panic(fmt.Errorf("parsing legal template: %w", err))
	}

	routes := []route.Route{
		route.NewHome(tplHome),
		route.NewSpryker(tplSpryker),
		route.NewLegal(tplLegal),
		route.NewStatic(*staticPath),
		route.NewRobots(),
	}

	s := server.NewServer(*port, routes)

	if err := s.Start(); err != nil {
		panic(err)
	}
}
