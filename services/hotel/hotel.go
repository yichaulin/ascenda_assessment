package hotel

type Hotel struct {
	ID                string    `json:"id"`
	Destination       string    `json:"destination"`
	Name              string    `json:"name"`
	Location          Location  `json:"location"`
	Description       string    `json:"description"`
	Amenities         Amenities `json:"amenities"`
	Images            Images    `json:"images"`
	BoodingConditions []string  `json:"booking_conditions"`
}

type Location struct {
	Lat     float32 `json:"lat"`
	Lng     float32 `json:"lng"`
	Address string  `json:"address"`
	City    string  `json:"city"`
	Contry  string  `json:"country"`
}

type Amenities struct {
	General []string `json:"general"`
	Room    []string `json:"room"`
	Others  []string `json:"others"`
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

func GetHotels(destination string, hotelIDs []string) (hotels []Hotel, err error) {
	hotels = []Hotel{
		{ID: "123", Destination: "456", Name: "789"},
		{ID: "qqq", Destination: "ooo", Name: "bbb"},
	}

	return hotels, nil
}
