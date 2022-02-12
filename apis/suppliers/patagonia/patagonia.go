package patagonia

import (
	"encoding/json"
	"strings"

	"ascenda_assessment/apis/resty"
	"ascenda_assessment/configs"
	amen "ascenda_assessment/utils/amenities"
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

func ParseAmenitiesToAmenityList(amenities []string) (general amen.AmenityList, room amen.AmenityList, others amen.AmenityList) {
	general = amen.AmenityList{}
	room = amen.AmenityList{}
	others = amen.AmenityList{}

	for _, a := range amenities {
		a = strings.TrimSpace(a)
		switch a {
		case Aircon:
			room.Add(amen.Aircon)
		case Tv:
			room.Add(amen.Tv)
		case CoffeeMachine:
			room.Add(amen.CoffeeMachine)
		case HairDryer:
			room.Add(amen.HairDryer)
		case Iron:
			room.Add(amen.Iron)
		case Tub:
			room.Add(amen.Bathtub)
		case Bar:
			general.Add(amen.Bar)
		case Kettle:
			room.Add(amen.Kettle)
		default:
			others.Add(a)
		}
	}

	return general, room, others
}
