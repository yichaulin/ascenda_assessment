package hotel

import (
	"ascenda_assessment/apis/suppliers/acme"
	"ascenda_assessment/apis/suppliers/paperflies"
	"ascenda_assessment/apis/suppliers/patagonia"
	"ascenda_assessment/logger"
	amen "ascenda_assessment/utils/amenities"
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
	Country *string  `json:"country"`
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

	paperflies, err := paperflies.GetData()
	if err != nil {
		logger.Error("Get Paperflies data failed", err)
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
		BoodingConditions: []string{},
		Images: Images{
			Rooms:     []ImageLink{},
			Site:      []ImageLink{},
			Amenities: []ImageLink{},
		},
	}

	return &hotel
}

func (hm hotelMap) toHotelSlice() []*Hotel {
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

		general, room, others := acme.ParseFacilitiesToAmenityList(d.Facilities)
		hotel.Amenities.GeneralList.Merge(general)
		hotel.Amenities.RoomList.Merge(room)
		hotel.Amenities.OthersList.Merge(others)
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

		general, room, others := patagonia.ParseAmenitiesToAmenityList(d.Amenities)
		hotel.Amenities.GeneralList.Merge(general)
		hotel.Amenities.RoomList.Merge(room)
		hotel.Amenities.OthersList.Merge(others)

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

func (hm hotelMap) mergePaperfliesData(paperfliesData []paperflies.PaperfliesData) {
	for _, d := range paperfliesData {
		hotel, ok := hm[d.HotelID]
		if !ok {
			hotel = newHotel()
			hm[d.HotelID] = hotel
		}

		hotel.ID = d.HotelID
		hotel.Destination = d.DestinationID
		hotel.Name = d.HotelName
		hotel.Location.Address = d.Location.Address
		hotel.Location.Country = d.Location.Country
		hotel.Description = d.Details
		hotel.BoodingConditions = d.BookingConditions

		general, room, others := paperflies.ParseAmenitiesToAmenityList(d.Amenities)
		hotel.Amenities.GeneralList.Merge(general)
		hotel.Amenities.RoomList.Merge(room)
		hotel.Amenities.OthersList.Merge(others)

		tmpImgs := make([]ImageLink, len(d.Images.Rooms))
		for i, roomImg := range d.Images.Rooms {
			tmpImgs[i] = ImageLink{
				Link:        roomImg.Link,
				Description: roomImg.Caption,
			}
		}
		hotel.Images.Rooms = append(hotel.Images.Rooms, tmpImgs...)

		tmpImgs = make([]ImageLink, len(d.Images.Site))
		for i, amenImg := range d.Images.Site {
			tmpImgs[i] = ImageLink{
				Link:        amenImg.Link,
				Description: amenImg.Caption,
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
