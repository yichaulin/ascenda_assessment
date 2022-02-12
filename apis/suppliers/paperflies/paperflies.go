package paperflies

import (
	"encoding/json"
	"strings"

	"ascenda_assessment/apis/resty"
	"ascenda_assessment/configs"
	amen "ascenda_assessment/utils/amenities"
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

func ParseAmenitiesToAmenityList(amenities Amenities) (general amen.AmenityList, room amen.AmenityList, others amen.AmenityList) {
	general = amen.AmenityList{}
	room = amen.AmenityList{}
	others = amen.AmenityList{}

	for _, a := range amenities.General {
		a = strings.TrimSpace(a)
		switch a {
		case GeneralOutdoorPool:
			general.Add(amen.OutdoorPool)
		case GeneralIndoorPool:
			general.Add(amen.IndoorPool)
		case GeneralBusinessCenter:
			general.Add(amen.BusinessCenter)
		case GeneralChildcare:
			general.Add(amen.Childcare)
		case GeneralParking:
			general.Add(amen.Parking)
		case GeneralBar:
			general.Add(amen.Bar)
		case GeneralDryCleaning:
			general.Add(amen.DryCleaning)
		case GeneralWifi:
			general.Add(amen.Wifi)
		case GeneralBreakfast:
			general.Add(amen.Breakfast)
		case GeneralConcierge:
			general.Add(amen.Concierge)
		default:
			others.Add(a)
		}
	}

	for _, a := range amenities.Room {
		a = strings.TrimSpace(a)
		switch a {
		case RoomTv:
			room.Add(amen.Tv)
		case RoomCoffeeMachine:
			room.Add(amen.CoffeeMachine)
		case RoomKettle:
			room.Add(amen.Kettle)
		case RoomHairDryer:
			room.Add(amen.HairDryer)
		case RoomIron:
			room.Add(amen.Iron)
		case RoomMinibar:
			room.Add(amen.Minibar)
		case RoomAircon:
			room.Add(amen.Aircon)
		case RoomBathtub:
			room.Add(amen.Bathtub)
		default:
			others.Add(a)
		}
	}

	return general, room, others
}
