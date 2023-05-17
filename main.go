package main

import (
	"ckilb/kilbtech/mail"
	"ckilb/kilbtech/route"
	"ckilb/kilbtech/tpl"
	_ "embed"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	port := flag.Int("port", 9090, "port")
	staticPath := flag.String("staticPath", "static", "staticPath")

	flag.Parse()

	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("loading dotenv: %w", err))
	}

	sender := mail.NewSender()

	routes := []route.Route{
		route.NewHome(),
		route.NewSpryker(),
		route.NewLegal(),
		route.NewRobots(),
		route.NewContact(sender),
	}

	engine := gin.Default()

	engine.StaticFS("/static", gin.Dir(*staticPath, false))

	renderer, err := tpl.NewRenderer()

	if err != nil {
		panic(err)
	}

	engine.HTMLRender = renderer

	for _, r := range routes {
		engine.Handle(r.Method(), r.Path(), r.Handler())
	}

	if err := engine.Run(fmt.Sprintf("127.0.0.1:%d", *port)); err != nil {
		panic(err)
	}
}
