package v1

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/erfidev/url_shortener/contract"
	"github.com/erfidev/url_shortener/dto"
	"github.com/erfidev/url_shortener/env"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func CreateUrl(i contract.UrlInteractor, app *env.AppEnv, ss *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		url := c.FormValue("url")

		var req dto.CreateShortUrlReq

		req.Url = url

		res, _ := i.CreateShortUrl(c.Context(), req)

		dtoData := dto.SummaryPageData{
			Domain: fmt.Sprintf("http://%s:%s", app.Domain, app.Port),
			SUrl:   res.SUrl,
			LUrl:   res.LUrl,
			Exp:    res.Exp,
			Msg:    res.Msg,
		}

		toJson, err := json.Marshal(dtoData)
		if err != nil {
			return c.Status(500).SendString("server error 500")
		}

		ss.Storage.Set("data", toJson, time.Minute)

		return c.Redirect("/summary")
	}
}
