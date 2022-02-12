package hotel

import (
	"ascenda_assessment/apis/suppliers/acme"
	"ascenda_assessment/apis/suppliers/paperflies"
	"ascenda_assessment/apis/suppliers/patagonia"
	"strings"
)

const (
	pool           = "pool"
	outdoorPool    = "outdoor pool"
	indoorPool     = "indoor pool"
	breakfast      = "breakfast"
	businessCenter = "business center"
	childcare      = "childcare"
	parking        = "parking"
	dryCleaning    = "dry cleaning"
	bar            = "bar"
	concierge      = "concierge"

	wifi          = "wifi"
	aircon        = "aircon"
	bathtub       = "bathtub"
	tv            = "tv"
	coffeeMachine = "coffee machine"
	hairDryer     = "hair dryer"
	iron          = "iron"
	kettle        = "kettle"
	minibar       = "minibar"
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
			room.Add(bathtub)
		case patagonia.Bar:
			genral.Add(bar)
		case patagonia.Kettle:
			room.Add(kettle)
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

func parsePaperfliesAmenities(amenities paperflies.Amenities) Amenities {
	general := amenityList{}
	room := amenityList{}
	others := amenityList{}

	for _, amen := range amenities.General {
		amen = strings.TrimSpace(amen)
		switch amen {
		case paperflies.GeneralOutdoorPool:
			general.Add(outdoorPool)
		case paperflies.GeneralIndoorPool:
			general.Add(indoorPool)
		case paperflies.GeneralBusinessCenter:
			general.Add(businessCenter)
		case paperflies.GeneralChildcare:
			general.Add(childcare)
		case paperflies.GeneralParking:
			general.Add(parking)
		case paperflies.GeneralBar:
			general.Add(bar)
		case paperflies.GeneralDryCleaning:
			general.Add(dryCleaning)
		case paperflies.GeneralWifi:
			general.Add(wifi)
		case paperflies.GeneralBreakfast:
			general.Add(breakfast)
		case paperflies.GeneralConcierge:
			general.Add(concierge)
		default:
			others.Add(amen)
		}
	}

	for _, amen := range amenities.Room {
		amen = strings.TrimSpace(amen)
		switch amen {
		case paperflies.RoomTv:
			room.Add(tv)
		case paperflies.RoomCoffeeMachine:
			room.Add(coffeeMachine)
		case paperflies.RoomKettle:
			room.Add(kettle)
		case paperflies.RoomHairDryer:
			room.Add(hairDryer)
		case paperflies.RoomIron:
			room.Add(iron)
		case paperflies.RoomMinibar:
			room.Add(minibar)
		case paperflies.RoomAircon:
			room.Add(aircon)
		case paperflies.RoomBathtub:
			room.Add(bathtub)
		default:
			others.Add(amen)
		}
	}

	return Amenities{
		GeneralList: general,
		RoomList:    room,
		OthersList:  others,
	}
}
