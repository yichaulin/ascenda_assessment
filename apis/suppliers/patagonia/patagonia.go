package patagonia

import (
	"encoding/json"
	"strings"
	"time"

	"ascenda_assessment/apis/client"
	"ascenda_assessment/configs"

	amen "ascenda_assessment/utils/amenities"
	"ascenda_assessment/utils/string_list"
	"ascenda_assessment/utils/supplier_data_filter"
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
	if cfg, ok := configs.Cfg.Suppliers[SupplierName]; ok && cfg.Timeout > 0 {
		ApiClient.SetTimeout(cfg.Timeout * time.Second)
	}
}

func (p PatagoniaData) GetHotelID() string {
	return p.ID
}

func (p PatagoniaData) GetDestinationID() uint64 {
	return p.DestinationID
}

func GetData(destination uint64, hotelIDs string_list.StringList) (interface{}, error) {
	patagoniaData := []PatagoniaData{}
	supplierConfig, ok := configs.Cfg.Suppliers[SupplierName]
	if !ok || !supplierConfig.Enabled {
		return patagoniaData, nil
	}

	resp, err := ApiClient.Get(supplierConfig.Url)
	if err != nil {
		return patagoniaData, err
	}

	tmp := []PatagoniaData{}
	json.Unmarshal(resp.Body(), &tmp)

	patagoniaData = make([]PatagoniaData, 0, len(tmp))
	for _, hotel := range tmp {
		if supplier_data_filter.IsMatchDestinationAndHotelID(hotel, destination, hotelIDs) {
			patagoniaData = append(patagoniaData, hotel)
		}
	}

	return patagoniaData, nil
}

func ParseAmenitiesToAmenityList(amenities []string) (general string_list.StringList, room string_list.StringList, others string_list.StringList) {
	general = string_list.StringList{}
	room = string_list.StringList{}
	others = string_list.StringList{}

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
