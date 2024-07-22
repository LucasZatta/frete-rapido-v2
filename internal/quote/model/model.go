package model

type Quote struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline int     `json:"deadline"`
	Price    float32 `json:"price"`
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
