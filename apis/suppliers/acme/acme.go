package acme

import (
	"ascenda_assessment/apis/resty"
	"ascenda_assessment/configs"
	"encoding/json"
)

type ACMEData struct {
	ID            string   `json:"Id"`
	DestinationID uint64   `json:"DestinationId"`
	Name          string   `json:"Name"`
	Latitude      *float32 `json:"Latitude"`
	Longitude     *float32 `json:"Longitude"`
	Address       *string  `json:"Address"`
	City          *string  `json:"City"`
	Country       *string  `json:"Country"`
	PostalCode    *string  `json:"PostalCode"`
	Description   *string  `json:"Description"`
	Facilities    []string `json:"Facilities"`
}

const (
	Pool           = "Pool"
	BusinessCenter = "BusinessCenter"
	WiFi           = "WiFi"
	DryCleaning    = "DryCleaning"
	Breakfast      = "Breakfast"
	Aircon         = "Aircon"
	BathTub        = "BathTub"
	Bar            = "Bar"
)

func GetData() (acmeData []ACMEData, err error) {
	url := configs.Cfg.Suppliers.ACME
	resp, err := resty.Get(url)
	if err != nil {
		return acmeData, err
	}

	json.Unmarshal(resp.Body(), &acmeData)

	return acmeData, nil
}
