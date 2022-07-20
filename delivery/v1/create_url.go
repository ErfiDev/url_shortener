package v1

import (
	"github.com/erfidev/url_shortener/contract"
	"github.com/erfidev/url_shortener/dto"
	"github.com/erfidev/url_shortener/env"
	"github.com/gofiber/fiber/v2"
)

func CreateUrl(i contract.UrlInteractor, app *env.AppEnv) fiber.Handler {
	return func(c *fiber.Ctx) error {
		url := c.FormValue("url")

		var req dto.CreateShortUrlReq

		req.Url = url

		res, _ := i.CreateShortUrl(c.Context(), req)

		return c.Render("summary_page", dto.SummaryPageData{
			Domain: app.Domain,
			SUrl:   res.SUrl,
			Exp:    res.Exp,
			LUrl:   res.LUrl,
			Msg:    res.Msg,
		}, "base_layout")
	}
}
