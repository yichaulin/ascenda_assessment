package hotel

import (
	"ascenda_assessment/apis/suppliers/acme"
	"strings"
)

const (
	pool           = "pool"
	breakfast      = "breakfast"
	businessCenter = "business center"
	wifi           = "wifi"
	dryCleaning    = "dry cleaning"
	aircon         = "aircon"
	bathtub        = "bathtub"
	bar            = "bar"
)

func parseACMEFacilities(facilities []string) Amenities {
	genral := amenityList{}
	room := amenityList{}
	others := amenityList{}

	for _, f := range facilities {
		amenity := strings.TrimSpace(f)
		switch amenity {
		case acme.Pool:
			genral.Add(pool)
		case acme.Breakfast:
			genral.Add(breakfast)
		case acme.BusinessCenter:
			genral.Add(businessCenter)
		case acme.WiFi:
			genral.Add(wifi)
		case acme.DryCleaning:
			genral.Add(dryCleaning)
		case acme.Aircon:
			room.Add(aircon)
		case acme.BathTub:
			room.Add(bathtub)
		case acme.Bar:
			genral.Add(bar)
		default:
			others.Add(amenity)
		}
	}

	return Amenities{
		GeneralList: genral,
		RoomList:    room,
		OthersList:  others,
	}
}
