package dto

type CreateShortUrlReq struct {
	Url string `json:"url"`
}

type GetUrlReq struct {
	SUrl string `json:"short_url"`
}
