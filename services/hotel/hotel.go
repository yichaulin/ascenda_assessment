package hotel

type Hotel struct {
	ID          string `json:"id"`
	Destination string `json:"destination"`
	Name        string `json:"Name"`
}

func GetHotels(destination string, hotelIDs []string) (hotels []Hotel, err error) {
	hotels = []Hotel{
		{ID: "123", Destination: "456", Name: "789"},
		{ID: "qqq", Destination: "ooo", Name: "bbb"},
	}

	return hotels, nil
}
