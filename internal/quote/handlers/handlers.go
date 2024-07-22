package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/model"
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
	path := os.Getenv("API_PATH")
	var requestBody model.QuoteReqBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// c.JSON(http.StatusAccepted, &requestBody)

	//organize and rename zipcode variables
	//treat possible errors
	//move this part of code to a utils package
	zipcode, _ := strconv.Atoi(requestBody.Recipient.Address.Zipcode)
	dispatcherZipcode, _ := strconv.Atoi(os.Getenv("DISPATCHER_ZIPCODE")) //utils -> regex to remove non number from zipcode
	volumesArray := make([]model.Volume, 0)
	fmt.Println("len: %i", len(requestBody.Volumes))

	for _, volume := range requestBody.Volumes {
		newVolum := model.Volume{
			Amount:        volume.Amount,
			Category:      strconv.Itoa(volume.Category),
			Height:        volume.Height,
			Width:         volume.Width,
			Length:        volume.Length,
			UnitaryPrice:  float64(volume.Price),
			UnitaryWeight: float64(volume.UnitaryWeight),
		}

		volumesArray = append(volumesArray, newVolum)
	}

	newDispatcher := model.Dispatcher{
		RegisteredNumber: os.Getenv("CNPJ"), //utils -> regex to remove non number from cnpj
		Zipcode:          dispatcherZipcode,
		Volumes:          volumesArray,
	}

	externalApiReqBody := &model.SimulationReqBody{
		Shipper: model.Shipper{
			RegisteredNumber: os.Getenv("CNPJ"),
			Token:            os.Getenv("API_TOKEN"),
			PlatformCode:     os.Getenv("PLATFORM_CODE"),
		},
		Recipient: model.Recipient{
			Type:    1,
			Country: "BRA", //consume api to track country by zipcode maybe?
			Zipcode: zipcode,
		},
		Dispatchers:    []model.Dispatcher{newDispatcher},
		SimulationType: []int{0},
		Returns: model.Returns{
			Composition:  false,
			Volumes:      true,
			AppliedRules: false,
		},
	}

	// c.JSON(http.StatusOK, &externalApiReqBody)

	// fmt.Printf("BUILT REQ BODY: %+v\n", externalApiReqBody)

	b, err := json.Marshal(externalApiReqBody)
	if err != nil {
		//treat marshal err
		fmt.Println(err)
		return
	}

	resp, err := http.Post(path, "application/json", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var simulationResponse model.SimulationResponse
	if err := json.Unmarshal(respBody, &simulationResponse); err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(resp.StatusCode, simulationResponse)

}
