package main

import (
	"ckilb/kilbtech/mail"
	"ckilb/kilbtech/route"
	"ckilb/kilbtech/server"
	_ "embed"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"text/template"
)

var (
	//go:embed template/contact.html
	contactTemplate string

	//go:embed template/layout.html
	layoutTemplate string

	//go:embed template/teaser.html
	teaserTemplate string

	//go:embed template/projects.html
	projectsTemplate string

	//go:embed template/project-scroller.html
	projectScrollerTemplate string

	//go:embed template/intro.html
	introTemplate string

	//go:embed template/teaser-intro.html
	teaserIntroTemplate string

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

	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("loading dotenv: %w", err))
	}

	sender := mail.NewSender()

	tplHome := template.New("home")
	tplSpryker := template.New("spryker")
	tplLegal := template.New("legal")

	if _, err := tplHome.Parse(projectScrollerTemplate + teaserIntroTemplate + introTemplate + contactTemplate + layoutTemplate + teaserTemplate + projectsTemplate + homeTemplate); err != nil {
		panic(fmt.Errorf("parsing home template: %w", err))
	}

	if _, err := tplSpryker.Parse(projectScrollerTemplate + teaserIntroTemplate + introTemplate + contactTemplate + layoutTemplate + teaserTemplate + projectsTemplate + sprykerTemplate); err != nil {
		panic(fmt.Errorf("parsing spryker template: %w", err))
	}

	if _, err := tplLegal.Parse(projectScrollerTemplate + teaserIntroTemplate + introTemplate + contactTemplate + layoutTemplate + teaserTemplate + projectsTemplate + legalTemplate); err != nil {
		panic(fmt.Errorf("parsing legal template: %w", err))
	}

	routes := []route.Route{
		route.NewHome(tplHome),
		route.NewSpryker(tplSpryker),
		route.NewLegal(tplLegal),
		route.NewStatic(*staticPath),
		route.NewRobots(),
		route.NewContact(sender),
	}

	s := server.NewServer(*port, routes)

	if err := s.Start(); err != nil {
		panic(err)
	}
}
