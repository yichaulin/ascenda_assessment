package patagonia

import (
	"encoding/json"
	"strings"
	"time"

	"ascenda_assessment/apis/client"
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
	SupplierName = "patagonia"

	Bar = "Bar"

	Aircon        = "Aircon"
	Tv            = "Tv"
	CoffeeMachine = "Coffee machine"
	Kettle        = "Kettle"
	HairDryer     = "Hair dryer"
	Iron          = "Iron"
	Tub           = "Tub"
)

var ApiClient client.Client

func init() {
	ApiClient = client.New()
	ApiClient.SetTimeout(5 * time.Second)
}

func GetData(destination uint64, hotelIDs map[string]struct{}) (patagoniaData []PatagoniaData, err error) {
	url := configs.Cfg.Suppliers.Patagonia
	resp, err := ApiClient.Get(url)
	if err != nil {
		return patagoniaData, err
	}

	tmp := []PatagoniaData{}
	json.Unmarshal(resp.Body(), &tmp)

	patagoniaData = make([]PatagoniaData, 0, len(tmp))
	for _, hotel := range tmp {
		_, matchHotelID := hotelIDs[hotel.ID]
		if (matchHotelID && hotel.DestinationID == destination) || (len(hotelIDs) == 0 && hotel.DestinationID == destination) {
			patagoniaData = append(patagoniaData, hotel)
		}
	}

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
