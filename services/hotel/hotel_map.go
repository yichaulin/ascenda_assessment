package hotel

import (
	"ascenda_assessment/apis/suppliers/acme"
	"ascenda_assessment/apis/suppliers/paperflies"
	"ascenda_assessment/apis/suppliers/patagonia"
	amen "ascenda_assessment/utils/amenities"
)

type hotelMap map[string]*Hotel

func (hm hotelMap) toHotelSlice() []*Hotel {
	hotels := make([]*Hotel, 0, len(hm))
	for _, h := range hm {
		amen.CleanGeneralListDuplicatedItem(h.Amenities.GeneralList)

		h.Amenities.General = h.Amenities.GeneralList.ToSlice()
		h.Amenities.Room = h.Amenities.RoomList.ToSlice()
		h.Amenities.Others = h.Amenities.OthersList.ToSlice()
		hotels = append(hotels, h)
	}

	return hotels
}

func (hm hotelMap) mergeSupplierData(data interface{}) {
	switch data.(type) {
	case []acme.ACMEData:
		hm.mergeACMEData(data.([]acme.ACMEData))
	case []patagonia.PatagoniaData:
		hm.mergePatagoniaData(data.([]patagonia.PatagoniaData))
	case []paperflies.PaperfliesData:
		hm.mergePaperfliesData(data.([]paperflies.PaperfliesData))
	default:
		return
	}
}

func (hm hotelMap) mergeACMEData(acmeData []acme.ACMEData) {
	supplier := acme.SupplierName

	for _, d := range acmeData {
		hotel, ok := hm[d.ID]
		if !ok {
			hotel = newHotel()
			hm[d.ID] = hotel
		}

		hotel.ID = d.ID
		hotel.Destination = d.DestinationID
		hotel.setNameWithPriority(d.Name, supplier)
		hotel.setDescriptionWithPriority(d.Description, supplier)
		hotel.Location.setLatLngWithPriority(d.Latitude, d.Longitude, supplier)
		hotel.Location.setAddressWithPriority(d.Address, supplier)
		hotel.Location.setCityWithPriority(d.City, supplier)

		general, room, others := acme.ParseFacilitiesToAmenityList(d.Facilities)
		hotel.Amenities.GeneralList.Merge(general)
		hotel.Amenities.RoomList.Merge(room)
		hotel.Amenities.OthersList.Merge(others)
	}
}

func (hm hotelMap) mergePatagoniaData(patagoniaData []patagonia.PatagoniaData) {
	supplier := patagonia.SupplierName

	for _, d := range patagoniaData {
		hotel, ok := hm[d.ID]
		if !ok {
			hotel = newHotel()
			hm[d.ID] = hotel
		}

		hotel.ID = d.ID
		hotel.Destination = d.DestinationID
		hotel.setNameWithPriority(d.Name, supplier)
		hotel.setDescriptionWithPriority(d.Info, supplier)
		hotel.Location.setLatLngWithPriority(d.Lat, d.Lng, supplier)
		hotel.Location.setAddressWithPriority(d.Address, supplier)

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
	supplier := paperflies.SupplierName

	for _, d := range paperfliesData {
		hotel, ok := hm[d.HotelID]
		if !ok {
			hotel = newHotel()
			hm[d.HotelID] = hotel
		}

		hotel.ID = d.HotelID
		hotel.Destination = d.DestinationID
		hotel.BookingConditions = d.BookingConditions
		hotel.setNameWithPriority(d.HotelName, supplier)
		hotel.setDescriptionWithPriority(d.Details, supplier)
		hotel.Location.setAddressWithPriority(d.Location.Address, supplier)
		hotel.Location.setCountryWithPriority(d.Location.Country, supplier)

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
