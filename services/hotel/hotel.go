package hotel

import (
	"fmt"
	"strconv"
	"strings"

	"ascenda_assessment/apis/suppliers/acme"
	"ascenda_assessment/apis/suppliers/paperflies"
	"ascenda_assessment/apis/suppliers/patagonia"
	"ascenda_assessment/configs"
	"ascenda_assessment/logger"

	amen "ascenda_assessment/utils/amenities"
)

type Hotel struct {
	ID                  string     `json:"id"`
	Destination         uint64     `json:"destination"`
	Name                string     `json:"name"`
	namePriority        uint       `json:"-"`
	Location            *Location  `json:"location"`
	Description         *string    `json:"description"`
	descriptionPriority uint       `json:"-"`
	Amenities           *Amenities `json:"amenities"`
	Images              Images     `json:"images"`
	BookingConditions   []string   `json:"booking_conditions"`
}

type Location struct {
	Lat             *float32 `json:"lat"`
	Lng             *float32 `json:"lng"`
	latLngPriority  uint     `json:"-"`
	Address         *string  `json:"address"`
	addressPriority uint     `json:"-"`
	City            *string  `json:"city"`
	cityPriority    uint     `json:"-"`
	Country         *string  `json:"country"`
	countryPriority uint     `json:"-"`
}

type Amenities struct {
	GeneralList amen.AmenityList `json:"-"`
	RoomList    amen.AmenityList `json:"-"`
	OthersList  amen.AmenityList `json:"-"`
	General     []string         `json:"general"`
	Room        []string         `json:"room"`
	Others      []string         `json:"others"`
}

type Images struct {
	Rooms     []ImageLink `json:"rooms"`
	Site      []ImageLink `json:"site"`
	Amenities []ImageLink `json:"amenities"`
}

type ImageLink struct {
	Link        string `json:"link"`
	Description string `json:"description"`
}

var InputInvalidError = fmt.Errorf("Invalid Input. Both Hotel IDs and Destination are empty")

func GetHotels(destination string, hotelIDs []string) (hotels []*Hotel, err error) {
	if len(destination) == 0 && len(hotelIDs) == 0 {
		return hotels, InputInvalidError
	}

	hm := make(hotelMap)

	destinationInt64, err := strconv.ParseUint(destination, 10, 64)
	if err != nil {
		return hotels, err
	}

	hotelIDList := map[string]struct{}{}
	for _, id := range hotelIDs {
		hotelIDList[id] = struct{}{}
	}

	acmeData, err := acme.GetData(destinationInt64, hotelIDList)
	if err != nil {
		logger.Error("Get ACME data failed.", err)
	}

	patagoniaData, err := patagonia.GetData(destinationInt64, hotelIDList)
	if err != nil {
		logger.Error("Get Patagonia data failed.", err)
	}

	paperflies, err := paperflies.GetData(destinationInt64, hotelIDList)
	if err != nil {
		logger.Error("Get Paperflies data failed.", err)
	}

	hm.mergeACMEData(acmeData)
	hm.mergePatagoniaData(patagoniaData)
	hm.mergePaperfliesData(paperflies)

	return hm.toHotelSlice(), nil
}

func newHotel() *Hotel {
	hotel := Hotel{
		Location: &Location{},
		Amenities: &Amenities{
			GeneralList: amen.AmenityList{},
			RoomList:    amen.AmenityList{},
			OthersList:  amen.AmenityList{},
		},
		BookingConditions: []string{},
		Images: Images{
			Rooms:     []ImageLink{},
			Site:      []ImageLink{},
			Amenities: []ImageLink{},
		},
	}

	return &hotel
}

func (h *Hotel) setNameWithPriority(name string, supplier string) {
	priority := configs.Cfg.SupplierDataPriorities.HotelName[supplier]
	if priority > h.namePriority {
		h.Name = name
		h.namePriority = priority
	}
}

func (h *Hotel) setDescriptionWithPriority(description *string, supplier string) {
	if description == nil {
		return
	}

	priority := configs.Cfg.SupplierDataPriorities.HotelDescription[supplier]
	if priority > h.descriptionPriority {
		desc := strings.TrimSpace(*description)
		h.Description = &desc
		h.descriptionPriority = priority
	}
}

func (l *Location) setCityWithPriority(city *string, supplier string) {
	if city == nil {
		return
	}

	priority := configs.Cfg.SupplierDataPriorities.HotelCity[supplier]
	if priority > l.cityPriority {
		c := strings.TrimSpace(*city)
		l.City = &c
		l.cityPriority = priority
	}
}

func (l *Location) setAddressWithPriority(address *string, supplier string) {
	if address == nil {
		return
	}

	priority := configs.Cfg.SupplierDataPriorities.HotelAddress[supplier]
	if priority > l.addressPriority {
		addr := strings.TrimSpace(*address)
		l.Address = &addr
		l.addressPriority = priority
	}
}

func (l *Location) setLatLngWithPriority(lat *float32, lng *float32, supplier string) {
	if lat == nil || lng == nil {
		return
	}

	priority := configs.Cfg.SupplierDataPriorities.HotelLatLng[supplier]
	if priority > l.latLngPriority {
		l.Lat = lat
		l.Lng = lng
		l.latLngPriority = priority
	}
}

func (l *Location) setCountryWithPriority(country *string, supplier string) {
	if country == nil {
		return
	}

	priority := configs.Cfg.SupplierDataPriorities.HotelCountry[supplier]
	if priority > l.countryPriority {
		c := strings.TrimSpace(*country)
		l.Country = &c
		l.countryPriority = priority
	}
}
