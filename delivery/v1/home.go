package v1

import (
	"fmt"
	"github.com/erfidev/url_shortener/dto"
	"github.com/erfidev/url_shortener/env"
	"github.com/gofiber/fiber/v2"
)

func Home(app *env.AppEnv) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("landing_page", dto.LandingPageData{
			Domain: fmt.Sprintf("https://%s:%s", app.Domain, app.Port),
		}, "base_layout")
	}
}
