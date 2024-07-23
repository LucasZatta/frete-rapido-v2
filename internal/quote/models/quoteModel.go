package models

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/lucaszatta/frete-rapido-v2/internal/util"
)

type Quote struct {
	ID        uint      `json:"-" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Service   string    `json:"service"`
	Deadline  int       `json:"deadline"`
	Price     float32   `json:"price"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type QuoteResponse struct {
	Carrier []Quote `json:"carrier"`
}

type QuoteReqBody struct {
	Recipient struct {
		Address struct {
			Zipcode string `json:"zipcode"`
		} `json:"address"`
	} `json:"recipient"`
	Volumes []struct {
		Category      int     `json:"category"`
		Amount        int     `json:"amount"`
		UnitaryWeight int     `json:"unitary_weight"`
		Price         int     `json:"price"`
		Sku           string  `json:"sku"`
		Height        float64 `json:"height"`
		Width         float64 `json:"width"`
		Length        float64 `json:"length"`
	} `json:"volumes"`
}

type LastQuotes struct {
	Name     string
	Count    int
	PriceSum float64
	PriceAvg float64
}

type LastQuotesResponse struct {
	LastQuotes      []LastQuotes
	ExpensiverQuote Quote
	CheapestQuote   Quote
}

func (qrb *QuoteReqBody) Validate() error {
	qrb.Recipient.Address.Zipcode = util.ClearString(qrb.Recipient.Address.Zipcode)
	if !util.ValidateZipcode(qrb.Recipient.Address.Zipcode) {
		return fmt.Errorf("invalid Zipcode: %s", qrb.Recipient.Address.Zipcode)
	}
	return nil
}

func (qrb *QuoteReqBody) BuildSimulationRequestBody() (*SimulationReqBody, error) {
	cnpj := util.ClearString(os.Getenv("CNPJ"))

	if !util.ValidateCNPJ(cnpj) {
		return &SimulationReqBody{}, errors.New("invalid cnpj")
	}

	err := qrb.Validate()
	if err != nil {
		return &SimulationReqBody{}, err
	}

	zipcode, err := strconv.Atoi(qrb.Recipient.Address.Zipcode)
	if err != nil {
		return &SimulationReqBody{}, err
	}

	dispatcherZipcode, err := strconv.Atoi(os.Getenv("DISPATCHER_ZIPCODE"))
	if err != nil {
		return &SimulationReqBody{}, err
	}

	volumesArray := make([]Volume, 0)

	for _, volume := range qrb.Volumes {
		newVolum := Volume{
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

	newDispatcher := Dispatcher{
		RegisteredNumber: os.Getenv("CNPJ"),
		Zipcode:          dispatcherZipcode,
		Volumes:          volumesArray,
	}

	return &SimulationReqBody{
		Shipper: Shipper{
			RegisteredNumber: os.Getenv("CNPJ"),
			Token:            os.Getenv("API_TOKEN"),
			PlatformCode:     os.Getenv("PLATFORM_CODE"),
		},
		Recipient: Recipient{
			Type:    1,
			Country: "BRA",
			Zipcode: zipcode,
		},
		Dispatchers:    []Dispatcher{newDispatcher},
		SimulationType: []int{0},
		Returns: Returns{
			Composition:  false,
			Volumes:      true,
			AppliedRules: false,
		},
	}, nil

}
