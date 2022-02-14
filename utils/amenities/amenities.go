package amenities

import "ascenda_assessment/utils/string_list"

const (
	Pool           = "pool"
	OutdoorPool    = "outdoor pool"
	IndoorPool     = "indoor pool"
	Breakfast      = "breakfast"
	BusinessCenter = "business center"
	Childcare      = "childcare"
	Parking        = "parking"
	DryCleaning    = "dry cleaning"
	Bar            = "bar"
	Concierge      = "concierge"

	Wifi          = "wifi"
	Aircon        = "aircon"
	Bathtub       = "bathtub"
	Tv            = "tv"
	CoffeeMachine = "coffee machine"
	HairDryer     = "hair dryer"
	Iron          = "iron"
	Kettle        = "kettle"
	Minibar       = "minibar"
)

func CleanGeneralListDuplicatedItem(list string_list.StringList) {
	_, hasOutdoorPool := list[OutdoorPool]
	_, hasIndoorPool := list[IndoorPool]
	if hasIndoorPool || hasOutdoorPool {
		list.Delete(Pool)
	}
}
