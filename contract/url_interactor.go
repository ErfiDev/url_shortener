package contract

import (
	"context"
	"github.com/erfidev/url_shortener/dto"
)

type UrlInteractor interface {
	CreateShortUrl(ctx context.Context, req dto.CreateShortUrlReq) (dto.CreateShortUrlRes, error)
	GetUrl(ctx context.Context, req dto.GetUrlReq) (dto.GetUrlRes, error)
}
