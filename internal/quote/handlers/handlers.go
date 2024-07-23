package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/models"
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
	lastQuotesParam := c.Query("last_quotes")

	lastQuotes, err := p.quoteService.GetLastQuotes(lastQuotesParam)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	expensiverQuote, cheapestQuote, err := p.quoteService.GetMaxMinQuotes()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	lastQuotesResponse := &models.LastQuotesResponse{
		LastQuotes:      *lastQuotes,
		ExpensiverQuote: *expensiverQuote,
		CheapestQuote:   *cheapestQuote,
	}

	c.JSON(http.StatusAccepted, lastQuotesResponse)

}

func (p *quoteHttp) SimulateQuoteHandler(c *gin.Context) {
	var requestBody models.QuoteReqBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	externalApiReqBody, err := requestBody.BuildSimulationRequestBody()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	b, err := json.Marshal(externalApiReqBody)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var simulationResponse models.SimulationResponse
	if gin.Mode() == gin.DebugMode {
		dummyResponse, err := os.Open("mockResponse.json")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		jsonParser := json.NewDecoder(dummyResponse)
		if err = jsonParser.Decode(&simulationResponse); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	} else {
		path := os.Getenv("API_PATH")
		resp, err := http.Post(path, "application/json", bytes.NewBuffer(b))
		if err != nil {
			c.AbortWithError(resp.StatusCode, err)
			return
		}

		defer resp.Body.Close()
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if err := json.Unmarshal(respBody, &simulationResponse); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	quotes := make([]models.Quote, 0)
	for _, dispatcher := range simulationResponse.Dispatchers {
		for _, offer := range dispatcher.Offers {
			newOffer := models.Quote{
				Name:     offer.Carrier.Name,
				Service:  offer.Service,
				Price:    float32(offer.FinalPrice),
				Deadline: offer.DeliveryTime.Days,
			}

			quotes = append(quotes, newOffer)
		}
	}

	created, err := p.quoteService.Create(&quotes)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusAccepted, created)

}
