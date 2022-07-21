package url

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/erfidev/url_shortener/dto"
	"github.com/erfidev/url_shortener/tools"
)

func (i Interact) CreateShortUrl(ctx context.Context, req dto.CreateShortUrlReq) (dto.CreateShortUrlRes, error) {
	_, err := url.ParseRequestURI(req.Url)
	if err != nil {
		return dto.CreateShortUrlRes{
			Msg: "Your url is not valid, change it",
		}, err
	}

	hash, err := tools.RandomHash(req.Url)
	if err != nil {
		return dto.CreateShortUrlRes{
			Msg: "Error from server, try again!",
		}, err
	}

	v, _ := i.DB.Get(ctx, hash)
	if v != "" {
		return dto.CreateShortUrlRes{
			Msg: "Error from server, try again!",
		}, err
	}

	err = i.DB.Set(ctx, hash, req.Url, time.Hour*2)
	if err != nil {
		return dto.CreateShortUrlRes{
			Msg: "Error from server, try again!",
		}, err
	}

	return dto.CreateShortUrlRes{
		SUrl: fmt.Sprintf("http://%s:%s/u/%s", i.Domain, i.Port, hash),
		LUrl: req.Url,
		Exp:  "Two Hour",
		Msg:  "Your short url successfuly generated!",
	}, nil
}
