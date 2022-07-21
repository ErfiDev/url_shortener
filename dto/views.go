package dto

type LandingPageData struct {
	Domain string
}

type SummaryPageData struct {
	Domain string `json:"domain"`
	SUrl   string `json:"short_url"`
	LUrl   string `json:"long_url"`
	Exp    string `json:"expiration"`
	Msg    string `json:"msg"`
}
