package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/service"
)

type QuoteHttp interface {
	GetQuote(c *gin.Context)
}

type quoteHttp struct {
	quoteService service.QuoteService
}

func New(quoteService service.QuoteService) *quoteHttp {
	return &quoteHttp{
		quoteService: quoteService,
	}
}

func (p *quoteHttp) GetQuote(c *gin.Context) {
	// ctx := r.Context()

	p.quoteService.GetQuote()

	//
	// id, err := quoteDecode.DecodeStringIDFromURI(r)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// product, err := p.productService.GetByID(ctx, id)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// encode.WriteJsonResponse(w, product, http.StatusOK)
}
