package hotel

import (
	"ascenda_assessment/apis/suppliers/acme"
	"ascenda_assessment/apis/suppliers/patagonia"
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
	tv             = "tv"
	coffeeMachine  = "coffee machine"
	hairDryer      = "hair dryer"
	iron           = "iron"
	tub            = "tub"
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

func parsePatagoniaAmenities(amenities []string) Amenities {
	genral := amenityList{}
	room := amenityList{}
	others := amenityList{}

	for _, a := range amenities {
		a = strings.TrimSpace(a)
		switch a {
		case patagonia.Aircon:
			room.Add(aircon)
		case patagonia.Tv:
			room.Add(tv)
		case patagonia.CoffeeMachine:
			room.Add(coffeeMachine)
		case patagonia.HairDryer:
			room.Add(hairDryer)
		case patagonia.Iron:
			room.Add(iron)
		case patagonia.Tub:
			room.Add(tub)
		case patagonia.Bar:
			genral.Add(bar)
		default:
			others.Add(a)
		}
	}

	return Amenities{
		GeneralList: genral,
		RoomList:    room,
		OthersList:  others,
	}
}
