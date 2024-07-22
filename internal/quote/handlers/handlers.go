package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/service"
)

type QuoteHttp interface {
	GetQuotes(c *gin.Context)
	SimulateQuote(c *gin.Context)
}

type quoteHttp struct {
	quoteService service.QuoteService
}

func New(quoteService service.QuoteService) *quoteHttp {
	return &quoteHttp{
		quoteService: quoteService,
	}
}

func (p *quoteHttp) GetQuotes(c *gin.Context) {
	lastQuotes := c.Query("last_quotes") // shortcut for c.Request.URL.Query().Get("lastname")

	fmt.Println(lastQuotes)
	if len(lastQuotes) == 0 {
		fmt.Println("quotes param is empty")
	}
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

func (p *quoteHttp) SimulateQuote(c *gin.Context) {
	p.quoteService.GetQuote()
}
