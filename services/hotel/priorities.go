package hotel

import (
	"ascenda_assessment/configs"
)

var (
	namePriorityConfig        = map[string]uint{}
	addressPriorityConfig     = map[string]uint{}
	descriptionPriorityConfig = map[string]uint{}
	countryPriorityConfig     = map[string]uint{}
	cityPriorityConfig        = map[string]uint{}
	latLngPriorityConfig      = map[string]uint{}
)

func init() {
	priorityCfg := configs.Cfg.SupplierDataPriorities

	namePriorityConfig = priorityCfg.HotelName
	addressPriorityConfig = priorityCfg.HotelAddress
	descriptionPriorityConfig = priorityCfg.HotelDescription
	countryPriorityConfig = priorityCfg.HotelCountry
	cityPriorityConfig = priorityCfg.HotelCity
	latLngPriorityConfig = priorityCfg.HotelLatLng
}
