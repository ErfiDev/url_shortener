package dto

type CreateShortUrlRes struct {
	SUrl string
	LUrl string
	Exp  string
	Msg  string
}

type GetUrlRes struct {
	LUrl string
}
