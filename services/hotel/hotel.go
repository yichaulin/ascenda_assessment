package hotel

import (
	"ascenda_assessment/apis/suppliers/acme"
	"ascenda_assessment/apis/suppliers/patagonia"
	"ascenda_assessment/logger"
	"strings"
)

type Hotel struct {
	ID                string     `json:"id"`
	Destination       uint64     `json:"destination"`
	Name              string     `json:"name"`
	Location          *Location  `json:"location"`
	Description       *string    `json:"description"`
	Amenities         *Amenities `json:"amenities"`
	Images            Images     `json:"images"`
	BoodingConditions []string   `json:"booking_conditions"`
}

type hotelMap map[string]*Hotel

type Location struct {
	Lat     *float32 `json:"lat"`
	Lng     *float32 `json:"lng"`
	Address *string  `json:"address"`
	City    *string  `json:"city"`
	Contry  *string  `json:"country"`
}

type Amenities struct {
	GeneralList amenityList `json:"-"`
	RoomList    amenityList `json:"-"`
	OthersList  amenityList `json:"-"`
	General     []string    `json:"general"`
	Room        []string    `json:"room"`
	Others      []string    `json:"others"`
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

func GetHotels(destination string, hotelIDs []string) (hotels []*Hotel, err error) {
	hm := make(hotelMap)

	acmeData, err := acme.GetData()
	if err != nil {
		logger.Error("Get ACME data failed", err)
	}

	patagoniaData, err := patagonia.GetData()
	if err != nil {
		logger.Error("Get Patagonia data failed", err)
	}

	hm.mergeACMEData(acmeData)
	hm.mergePatagoniaData(patagoniaData)

	return hm.ToHotelSlice(), nil
}

func newHotel() *Hotel {
	hotel := Hotel{
		Location: &Location{},
		Amenities: &Amenities{
			GeneralList: amenityList{},
			RoomList:    amenityList{},
			OthersList:  amenityList{},
		},
		BoodingConditions: []string{},
		Images: Images{
			Rooms:     []ImageLink{},
			Site:      []ImageLink{},
			Amenities: []ImageLink{},
		},
	}

	return &hotel
}

func (hm hotelMap) ToHotelSlice() []*Hotel {
	hotels := make([]*Hotel, 0, len(hm))
	for _, h := range hm {
		h.Amenities.ListToStringSlice()
		hotels = append(hotels, h)
	}

	return hotels
}

func (hm hotelMap) mergeACMEData(acmeData []acme.ACMEData) {
	for _, d := range acmeData {
		hotel, ok := hm[d.ID]
		if !ok {
			hotel = newHotel()
			hm[d.ID] = hotel
		}

		hotel.ID = d.ID
		hotel.Destination = d.DestinationID
		hotel.Name = d.Name
		hotel.Location.Lat = d.Latitude
		hotel.Location.Lng = d.Longitude
		if d.Address != nil {
			addr := strings.TrimSpace(*d.Address)
			hotel.Location.Address = &addr
		}
		if d.Description != nil {
			desc := strings.TrimSpace(*d.Description)
			hotel.Description = &desc
		}
		if d.City != nil {
			city := strings.TrimSpace(*d.City)
			hotel.Location.City = &city
		}

		amenities := parseACMEFacilities(d.Facilities)
		hotel.Amenities.GeneralList.Merge(amenities.GeneralList)
		hotel.Amenities.RoomList.Merge(amenities.RoomList)
		hotel.Amenities.OthersList.Merge(amenities.OthersList)

	}
}

func (hm hotelMap) mergePatagoniaData(patagoniaData []patagonia.PatagoniaData) {
	for _, d := range patagoniaData {
		hotel, ok := hm[d.ID]
		if !ok {
			hotel = newHotel()
			hm[d.ID] = hotel
		}

		hotel.ID = d.ID
		hotel.Destination = d.DestinationID
		hotel.Name = d.Name
		hotel.Location.Lat = d.Lat
		hotel.Location.Lng = d.Lng
		if d.Info != nil {
			desc := strings.TrimSpace(*d.Info)
			hotel.Description = &desc
		}
		if d.Address != nil {
			addr := strings.TrimSpace(*d.Address)
			hotel.Location.Address = &addr
		}

		amenities := parsePatagoniaAmenities(d.Amenities)
		hotel.Amenities.GeneralList.Merge(amenities.GeneralList)
		hotel.Amenities.RoomList.Merge(amenities.RoomList)
		hotel.Amenities.OthersList.Merge(amenities.OthersList)

		tmpImgs := make([]ImageLink, len(d.Images.Rooms))
		for i, roomImg := range d.Images.Rooms {
			tmpImgs[i] = ImageLink{
				Link:        roomImg.Url,
				Description: roomImg.Description,
			}
		}
		hotel.Images.Rooms = append(hotel.Images.Rooms, tmpImgs...)

		tmpImgs = make([]ImageLink, len(d.Images.Amenities))
		for i, amenImg := range d.Images.Amenities {
			tmpImgs[i] = ImageLink{
				Link:        amenImg.Url,
				Description: amenImg.Description,
			}
		}
		hotel.Images.Amenities = append(hotel.Images.Amenities, tmpImgs...)
	}
}

func (a *Amenities) ListToStringSlice() {
	a.General = a.GeneralList.ToStringSlice()
	a.Room = a.RoomList.ToStringSlice()
	a.Others = a.OthersList.ToStringSlice()
}
