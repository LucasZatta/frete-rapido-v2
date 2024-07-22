package models

type Shipper struct {
	RegisteredNumber string `json:"registered_number"`
	Token            string `json:"token"`
	PlatformCode     string `json:"platform_code"`
}
type Recipient struct {
	Type             int    `json:"type"`
	RegisteredNumber string `json:"registered_number"`
	StateInscription string `json:"state_inscription"`
	Country          string `json:"country"`
	Zipcode          int    `json:"zipcode"`
}

type Volume struct {
	Amount        int     `json:"amount"`
	AmountVolumes int     `json:"amount_volumes"`
	Category      string  `json:"category"`
	Sku           string  `json:"sku"`
	Tag           string  `json:"tag"`
	Description   string  `json:"description"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
	UnitaryPrice  float64 `json:"unitary_price"`
	UnitaryWeight float64 `json:"unitary_weight"`
	Consolidate   bool    `json:"consolidate"`
	Overlaid      bool    `json:"overlaid"`
	Rotate        bool    `json:"rotate"`
}

type Dispatcher struct {
	RegisteredNumber string   `json:"registered_number"`
	Zipcode          int      `json:"zipcode"`
	TotalPrice       float64  `json:"total_price"`
	Volumes          []Volume `json:"volumes"`
}

type Returns struct {
	Composition  bool `json:"composition"`
	Volumes      bool `json:"volumes"`
	AppliedRules bool `json:"applied_rules"`
}

type SimulationReqBody struct {
	Shipper        Shipper      `json:"shipper"`
	Recipient      Recipient    `json:"recipient"`
	Dispatchers    []Dispatcher `json:"dispatchers"`
	Channel        string       `json:"channel"`
	Filter         int          `json:"filter"`
	Limit          int          `json:"limit"`
	Identification string       `json:"identification"`
	Reverse        bool         `json:"reverse"`
	SimulationType []int        `json:"simulation_type"`
	Returns        Returns      `json:"returns"`
}

type Offer struct {
	Offer          int `json:"offer"`
	SimulationType int `json:"simulation_type"`
	Carrier        struct {
		Reference        int    `json:"reference"`
		Name             string `json:"name"`
		RegisteredNumber string `json:"registered_number"`
		StateInscription string `json:"state_inscription"`
		Logo             string `json:"logo"`
	} `json:"carrier"`
	Service            string `json:"service"`
	ServiceCode        string `json:"service_code"`
	ServiceDescription string `json:"service_description"`
	DeliveryTime       struct {
		Days          int    `json:"days"`
		Hours         int    `json:"hours"`
		Minutes       int    `json:"minutes"`
		EstimatedDate string `json:"estimated_date"`
	} `json:"delivery_time"`
	Expiration string  `json:"expiration"`
	CostPrice  float64 `json:"cost_price"`
	FinalPrice float64 `json:"final_price"`
	Weights    struct {
		Real  float64 `json:"real"`
		Cubed float64 `json:"cubed"`
		Used  float64 `json:"used"`
	} `json:"weights"`
	Composition struct {
		FreightWeight       float64 `json:"freight_weight"`
		FreightWeightExcess float64 `json:"freight_weight_excess"`
		FreightWeightVolume float64 `json:"freight_weight_volume"`
		FreightVolume       float64 `json:"freight_volume"`
		FreightMinimum      float64 `json:"freight_minimum"`
		FreightInvoice      float64 `json:"freight_invoice"`
		SubTotal1           struct {
			Daily           int     `json:"daily"`
			Collect         int     `json:"collect"`
			Dispatch        int     `json:"dispatch"`
			Delivery        int     `json:"delivery"`
			Ferry           int     `json:"ferry"`
			Suframa         int     `json:"suframa"`
			Tas             int     `json:"tas"`
			SecCat          int     `json:"sec_cat"`
			Dat             int     `json:"dat"`
			AdValorem       float64 `json:"ad_valorem"`
			Ademe           int     `json:"ademe"`
			Gris            float64 `json:"gris"`
			Emex            int     `json:"emex"`
			Interior        int     `json:"interior"`
			Capatazia       int     `json:"capatazia"`
			River           int     `json:"river"`
			RiverInsurance  int     `json:"river_insurance"`
			Toll            int     `json:"toll"`
			Other           int     `json:"other"`
			OtherPerProduct int     `json:"other_per_product"`
		} `json:"sub_total1"`
		SubTotal2 struct {
			Trt        int `json:"trt"`
			Tda        int `json:"tda"`
			Tde        int `json:"tde"`
			Scheduling int `json:"scheduling"`
		} `json:"sub_total2"`
		SubTotal3 struct {
			Icms float64 `json:"icms"`
		} `json:"sub_total3"`
	} `json:"composition"`
	OriginalDeliveryTime struct {
		Days          int    `json:"days"`
		Hours         int    `json:"hours"`
		Minutes       int    `json:"minutes"`
		EstimatedDate string `json:"estimated_date"`
	} `json:"original_delivery_time"`
	Identifier                   string `json:"identifier"`
	DeliveryNote                 string `json:"delivery_note"`
	HomeDelivery                 bool   `json:"home_delivery"`
	CarrierNeedsToReturnToSender bool   `json:"carrier_needs_to_return_to_sender"`
	Modal                        string `json:"modal"`
	Esg                          struct {
		Co2EmissionEstimate   float64 `json:"co2_emission_estimate"`
		Co2NeutralizationCost float64 `json:"co2_neutralization_cost"`
	} `json:"esg"`
}

type SimulationResponse struct {
	Dispatchers []struct {
		ID                         string  `json:"id"`
		RequestID                  string  `json:"request_id"`
		RegisteredNumberShipper    string  `json:"registered_number_shipper"`
		RegisteredNumberDispatcher string  `json:"registered_number_dispatcher"`
		ZipcodeOrigin              int     `json:"zipcode_origin"`
		Offers                     []Offer `json:"offers"`
		Volumes                    []struct {
			Category      string        `json:"category"`
			Sku           string        `json:"sku"`
			Tag           string        `json:"tag"`
			Description   string        `json:"description"`
			Amount        int           `json:"amount"`
			Width         float64       `json:"width"`
			Height        float64       `json:"height"`
			Length        float64       `json:"length"`
			UnitaryWeight float64       `json:"unitary_weight"`
			UnitaryPrice  float64       `json:"unitary_price"`
			AmountVolumes float64       `json:"amount_volumes"`
			Consolidate   bool          `json:"consolidate"`
			Overlaid      bool          `json:"overlaid"`
			Rotate        bool          `json:"rotate"`
			Items         []interface{} `json:"items"`
		} `json:"volumes"`
	} `json:"dispatchers"`
}
