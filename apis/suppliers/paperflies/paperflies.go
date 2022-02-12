package paperflies

import (
	"ascenda_assessment/apis/resty"
	"ascenda_assessment/configs"
	"encoding/json"
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
	GeneralOutdoorPool    = "outdoor pool"
	GeneralIndoorPool     = "indoor pool"
	GeneralBusinessCenter = "business center"
	GeneralChildcare      = "childcare"
	GeneralParking        = "parking"
	GeneralBar            = "bar"
	GeneralDryCleaning    = "dry cleaning"
	GeneralWifi           = "wifi"
	GeneralBreakfast      = "breakfast"
	GeneralConcierge      = "concierge"

	RoomTv            = "tv"
	RoomCoffeeMachine = "coffee machine"
	RoomKettle        = "kettle"
	RoomHairDryer     = "hair dryer"
	RoomIron          = "iron"
	RoomMinibar       = "minibar"
	RoomBathtub       = "bathtub"
	RoomAircon        = "aircon"
)

func GetData() (paperfliesData []PaperfliesData, err error) {
	url := configs.Cfg.Suppliers.Paperflies
	resp, err := resty.Get(url)
	if err != nil {
		return paperfliesData, err
	}

	json.Unmarshal(resp.Body(), &paperfliesData)

	return paperfliesData, nil
}
