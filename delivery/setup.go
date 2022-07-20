package delivery

import (
	"github.com/erfidev/url_shortener/contract"
	"github.com/erfidev/url_shortener/env"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Setup(db contract.DbContract, app *env.AppEnv) {
	r := fiber.New()

	log.Printf("router initialized, listening port: %s\n", app.Port)
	err := r.Listen(app.Port)
	if err != nil {
		log.Fatalf("Error on running router: %s", err)
	}
}
