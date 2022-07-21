package v1

import (
	"encoding/json"

	"github.com/erfidev/url_shortener/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Summary(ss *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		v, err := ss.Storage.Get("data")
		if err != nil {
			return c.Status(500).SendString("server error 500")
		}

		var decode dto.SummaryPageData
		err = json.Unmarshal(v, &decode)
		if err != nil {
			return c.Status(500).SendString("server error 500")
		}

		return c.Render("summary_page", decode)
	}
}
