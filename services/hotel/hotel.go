package hotel

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	supplier "ascenda_assessment/apis/suppliers"
	"ascenda_assessment/configs"
	"ascenda_assessment/logger"

	"ascenda_assessment/utils/string_list"
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
	GeneralList string_list.StringList `json:"-"`
	RoomList    string_list.StringList `json:"-"`
	OthersList  string_list.StringList `json:"-"`
	General     []string               `json:"general"`
	Room        []string               `json:"room"`
	Others      []string               `json:"others"`
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

	destinationInt64, _ := strconv.ParseUint(destination, 10, 64)
	hotelIDList := string_list.New(hotelIDs...)

	wg := new(sync.WaitGroup)
	suppliers := supplier.GetAllSupplierNames()
	wg.Add(len(suppliers))
	for _, s := range suppliers {
		go func(wg *sync.WaitGroup, spl string) {
			defer wg.Done()
			data, e := supplier.GetData(spl, destinationInt64, hotelIDList)
			if e != nil {
				logger.Error(fmt.Sprintf("Get %s data failed.", spl), e)
				return
			}
			hm.mergeSupplierData(data)
		}(wg, s)
	}
	wg.Wait()

	return hm.toHotelSlice(), nil
}

func newHotel() *Hotel {
	hotel := Hotel{
		Location: &Location{},
		Amenities: &Amenities{
			GeneralList: string_list.StringList{},
			RoomList:    string_list.StringList{},
			OthersList:  string_list.StringList{},
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
