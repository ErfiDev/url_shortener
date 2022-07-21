package delivery

import (
	"log"

	"github.com/erfidev/url_shortener/contract"
	v1 "github.com/erfidev/url_shortener/delivery/v1"
	"github.com/erfidev/url_shortener/env"
	"github.com/erfidev/url_shortener/interactors/url"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func Setup(db contract.DbContract, app *env.AppEnv) {
	viewEngine := html.New("./views", ".html")

	r := fiber.New(fiber.Config{
		Views:       viewEngine,
		ViewsLayout: "base_layout",
	})

	// static files
	r.Static("/static", "./static")

	// initializing interactors
	urlInteract := url.New(db, app.Domain)

	// pages
	r.Get("/", v1.Home(app))
	r.Get("/u/:url", v1.GetUrl(urlInteract))

	// post methods
	r.Post("/api/v1/create_url", v1.CreateUrl(urlInteract, app))

	log.Printf("router initialized, listening port: %s\n", app.Port)
	err := r.Listen(":" + app.Port)
	if err != nil {
		log.Fatalf("Error on running router: %s", err)
	}
}
