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
