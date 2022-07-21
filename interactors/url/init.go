package url

import "github.com/erfidev/url_shortener/contract"

type Interact struct {
	DB     contract.DbContract
	Domain string
	Port   string
}

func New(db contract.DbContract, dom, port string) Interact {
	return Interact{
		DB:     db,
		Domain: dom,
		Port:   port,
	}
}
