package v1

import (
	"github.com/erfidev/url_shortener/contract"
	"github.com/erfidev/url_shortener/dto"
	"github.com/gofiber/fiber/v2"
)

func GetUrl(i contract.UrlInteractor) fiber.Handler {
	return func(c *fiber.Ctx) error {
		url := c.Params("url")
		if url == "" {
			return c.Status(404).SendString("NOT FOUND!")
		}

		req := dto.GetUrlReq{
			SUrl: url,
		}

		res, err := i.GetUrl(c.Context(), req)
		if err != nil {
			return c.Status(404).SendString("NOT FOUND!")
		}

		return c.Redirect(res.LUrl, fiber.StatusPermanentRedirect)
	}
}
