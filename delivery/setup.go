package delivery

import (
	"github.com/erfidev/url_shortener/contract"
	"github.com/erfidev/url_shortener/env"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"log"
)

func Setup(db contract.DbContract, app *env.AppEnv) {
	viewEngine := html.New("./views", ".gohtml")

	r := fiber.New(fiber.Config{
		Views: viewEngine,
	})

	// home
	r.Get("/", )

	log.Printf("router initialized, listening port: %s\n", app.Port)
	err := r.Listen(app.Port)
	if err != nil {
		log.Fatalf("Error on running router: %s", err)
	}
}
