package http

import "github.com/lucaszatta/frete-rapido-v2/internal/quote/service"

type QuoteHttp interface {
}

type quoteHttp struct {
	quoteService service.QuoteService
}

func New(quoteService service.QuoteService) *quoteHttp {
	return &quoteHttp{
		quoteService: quoteService,
	}
}
