package url

import (
	"context"

	"github.com/erfidev/url_shortener/dto"
)

func (i Interact) GetUrl(ctx context.Context, req dto.GetUrlReq) (dto.GetUrlRes, error) {
	v, err := i.DB.Get(ctx, req.SUrl)
	if err != nil {
		return dto.GetUrlRes{}, err
	}

	if v == "" {
		return dto.GetUrlRes{}, err
	}

	return dto.GetUrlRes{
		LUrl: v,
	}, nil
}
