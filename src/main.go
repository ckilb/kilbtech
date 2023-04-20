package main

import (
	"ckilb/kilbtech/mail"
	"ckilb/kilbtech/route"
	"ckilb/kilbtech/server"
	"ckilb/kilbtech/tpl"
	_ "embed"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	port := flag.Int("port", 9090, "port")
	staticPath := flag.String("staticPath", "../static", "staticPath")

	flag.Parse()

	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("loading dotenv: %w", err))
	}

	sender := mail.NewSender()
	renderer := tpl.NewRenderer()

	routes := []route.Route{
		route.NewHome(renderer),
		route.NewSpryker(renderer),
		route.NewLegal(renderer),
		route.NewStatic(*staticPath),
		route.NewRobots(),
		route.NewContact(sender),
	}

	s := server.NewServer(*port, routes)

	if err := s.Start(); err != nil {
		panic(err)
	}
}
