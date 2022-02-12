package patagonia

import (
	"ascenda_assessment/apis/resty"
	"ascenda_assessment/configs"
	"encoding/json"
)

type PatagoniaData struct {
	ID            string          `json:"id"`
	DestinationID uint64          `json:"destination"`
	Name          string          `json:"name"`
	Lat           *float32        `json:"lat"`
	Lng           *float32        `json:"lng"`
	Address       *string         `json:"address"`
	Info          *string         `json:"info"`
	Amenities     []string        `json:"amenities"`
	Images        PatagoniaImages `json:"images"`
}

type PatagoniaImages struct {
	Rooms     []Image `json:"rooms"`
	Amenities []Image `json:"amenities"`
}

type Image struct {
	Url         string `json:"url"`
	Description string `json:"description"`
}

const (
	Aircon        = "Aircon"
	Tv            = "Tv"
	CoffeeMachine = "Coffee machine"
	Kettle        = "Kettle"
	HairDryer     = "Hair dryer"
	Iron          = "Iron"
	Tub           = "Tub"
	Bar           = "Bar"
)

func GetData() (patagoniaData []PatagoniaData, err error) {
	url := configs.Cfg.Suppliers.Patagonia
	resp, err := resty.Get(url)
	if err != nil {
		return patagoniaData, err
	}

	json.Unmarshal(resp.Body(), &patagoniaData)

	return patagoniaData, nil
}
