package paperflies

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

type PaperfliesData struct {
	HotelID           string    `json:"hotel_id"`
	DestinationID     uint64    `json:"destination_id"`
	HotelName         string    `json:"hotel_name"`
	Location          Location  `json:"location"`
	Details           *string   `json:"details"`
	Amenities         Amenities `json:"amenities"`
	Images            Images    `json:"images"`
	BookingConditions []string  `json:"booking_conditions"`
}

type Location struct {
	Address *string `json:"address"`
	Country *string `json:"country"`
}

type Amenities struct {
	General []string `json:"general"`
	Room    []string `json:"room"`
}

type Images struct {
	Rooms []Image `json:"rooms"`
	Site  []Image `json:"site"`
}

type Image struct {
	Link    string `json:"link"`
	Caption string `json:"caption"`
}

const (
	SupplierName = "paperflies"

	OutdoorPool    = "outdoor pool"
	IndoorPool     = "indoor pool"
	BusinessCenter = "business center"
	Childcare      = "childcare"
	Parking        = "parking"
	Bar            = "bar"
	DryCleaning    = "dry cleaning"
	Wifi           = "wifi"
	Breakfast      = "breakfast"
	Concierge      = "concierge"

	Tv            = "tv"
	CoffeeMachine = "coffee machine"
	Kettle        = "kettle"
	HairDryer     = "hair dryer"
	Iron          = "iron"
	Minibar       = "minibar"
	Bathtub       = "bathtub"
	Aircon        = "aircon"
)

var ApiClient client.Client

func init() {
	ApiClient = client.New()
	if cfg, ok := configs.Cfg.Suppliers[SupplierName]; ok && cfg.Timeout > 0 {
		ApiClient.SetTimeout(cfg.Timeout * time.Second)
	}
}

func (p PaperfliesData) GetHotelID() string {
	return p.HotelID
}

func (p PaperfliesData) GetDestinationID() uint64 {
	return p.DestinationID
}

func GetData(destination uint64, hotelIDs string_list.StringList) (interface{}, error) {
	supplierConfig, ok := configs.Cfg.Suppliers[SupplierName]
	paperfliesData := []PaperfliesData{}

	if !ok || !supplierConfig.Enabled {
		return paperfliesData, nil
	}

	resp, err := ApiClient.Get(supplierConfig.Url)
	if err != nil {
		return paperfliesData, err
	}

	tmp := []PaperfliesData{}
	json.Unmarshal(resp.Body(), &tmp)

	paperfliesData = make([]PaperfliesData, 0, len(tmp))
	for _, hotel := range tmp {
		if supplier_data_filter.IsMatchDestinationAndHotelID(hotel, destination, hotelIDs) {
			paperfliesData = append(paperfliesData, hotel)
		}
	}

	return paperfliesData, nil
}

func ParseAmenitiesToAmenityList(amenities Amenities) (general string_list.StringList, room string_list.StringList, others string_list.StringList) {
	general = string_list.StringList{}
	room = string_list.StringList{}
	others = string_list.StringList{}

	for _, a := range amenities.General {
		a = strings.TrimSpace(a)
		switch a {
		case OutdoorPool:
			general.Add(amen.OutdoorPool)
		case IndoorPool:
			general.Add(amen.IndoorPool)
		case BusinessCenter:
			general.Add(amen.BusinessCenter)
		case Childcare:
			general.Add(amen.Childcare)
		case Parking:
			general.Add(amen.Parking)
		case Bar:
			general.Add(amen.Bar)
		case DryCleaning:
			general.Add(amen.DryCleaning)
		case Wifi:
			general.Add(amen.Wifi)
		case Breakfast:
			general.Add(amen.Breakfast)
		case Concierge:
			general.Add(amen.Concierge)
		default:
			others.Add(a)
		}
	}

	for _, a := range amenities.Room {
		a = strings.TrimSpace(a)
		switch a {
		case Tv:
			room.Add(amen.Tv)
		case CoffeeMachine:
			room.Add(amen.CoffeeMachine)
		case Kettle:
			room.Add(amen.Kettle)
		case HairDryer:
			room.Add(amen.HairDryer)
		case Iron:
			room.Add(amen.Iron)
		case Minibar:
			room.Add(amen.Minibar)
		case Aircon:
			room.Add(amen.Aircon)
		case Bathtub:
			room.Add(amen.Bathtub)
		default:
			others.Add(a)
		}
	}

	return general, room, others
}
