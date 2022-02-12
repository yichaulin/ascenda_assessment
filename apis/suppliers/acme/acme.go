package acme

import (
	"encoding/json"
	"strings"

	"ascenda_assessment/apis/resty"
	"ascenda_assessment/configs"

	amen "ascenda_assessment/utils/amenities"
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

func GetData(destination uint64, hotelIDs map[string]struct{}) (acmeData []ACMEData, err error) {
	url := configs.Cfg.Suppliers.ACME
	resp, err := resty.Get(url)
	if err != nil {
		return acmeData, err
	}

	tmp := []ACMEData{}
	json.Unmarshal(resp.Body(), &tmp)

	acmeData = make([]ACMEData, 0, len(tmp))
	for _, hotel := range tmp {
		_, matchHotelID := hotelIDs[hotel.ID]
		if (matchHotelID && hotel.DestinationID == destination) || (len(hotelIDs) == 0 && hotel.DestinationID == destination) {
			acmeData = append(acmeData, hotel)
		}
	}

	return acmeData, nil
}

func ParseFacilitiesToAmenityList(facilities []string) (general amen.AmenityList, room amen.AmenityList, others amen.AmenityList) {
	general = amen.AmenityList{}
	room = amen.AmenityList{}
	others = amen.AmenityList{}

	for _, f := range facilities {
		f := strings.TrimSpace(f)
		switch f {
		case Pool:
			general.Add(amen.Pool)
		case Breakfast:
			general.Add(amen.Breakfast)
		case BusinessCenter:
			general.Add(amen.BusinessCenter)
		case WiFi:
			general.Add(amen.Wifi)
		case DryCleaning:
			general.Add(amen.DryCleaning)
		case Aircon:
			room.Add(amen.Aircon)
		case BathTub:
			room.Add(amen.Bathtub)
		case Bar:
			general.Add(amen.Bar)
		default:
			others.Add(f)
		}
	}

	return general, room, others
}
