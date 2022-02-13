package acme

import (
	"encoding/json"
	"strings"
	"time"

	"ascenda_assessment/apis/client"
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
	SupplierName = "acme"

	Pool           = "Pool"
	BusinessCenter = "BusinessCenter"
	DryCleaning    = "DryCleaning"
	Breakfast      = "Breakfast"
	Bar            = "Bar"
	WiFi           = "WiFi"

	Aircon  = "Aircon"
	BathTub = "BathTub"
)

var ApiClient client.Client

func init() {
	ApiClient = client.New()
	ApiClient.SetTimeout(5 * time.Second)
}

func GetData(destination uint64, hotelIDs map[string]struct{}) (acmeData []ACMEData, err error) {
	url := configs.Cfg.Suppliers.ACME
	resp, err := ApiClient.Get(url)
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
