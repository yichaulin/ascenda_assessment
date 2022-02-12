package amenities

type AmenityList map[string]struct{}

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

func CleanGeneralListDuplicatedItem(list AmenityList) {
	_, hasOutdoorPool := list[OutdoorPool]
	_, hasIndoorPool := list[IndoorPool]
	if hasIndoorPool || hasOutdoorPool {
		delete(list, Pool)
	}
}

func (l AmenityList) ToStringSlice() []string {
	s := make([]string, 0, len(l))
	for k := range l {
		s = append(s, k)
	}

	return s
}

func (l AmenityList) Add(amens ...string) {
	for _, amen := range amens {
		l[amen] = struct{}{}
	}
}

func (l AmenityList) Merge(from AmenityList) {
	for k := range from {
		l[k] = struct{}{}
	}
}
