package paperflies_test

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	"ascenda_assessment/apis/suppliers/paperflies"
	"ascenda_assessment/configs"
	"ascenda_assessment/tests/mock"
	amen "ascenda_assessment/utils/amenities"
)

func TestGetData(t *testing.T) {
	ass := assert.New(t)

	url := configs.Cfg.Suppliers.Paperflies
	mockDataPath := "./mockData.json"
	mockBody, err := os.ReadFile(mockDataPath)
	ass.Nil(err, fmt.Sprintf("Read Paperflies mock data file failed. Path: %s", mockDataPath))

	mock.MockAPI(paperflies.ApiClient, "GET", url, string(mockBody), http.StatusOK)
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		destination       uint64
		hotelIDs          []string
		expectHotelCounts int
	}{
		{
			destination:       5432,
			hotelIDs:          []string{"SjyX", "iJhz", "InvalidID"},
			expectHotelCounts: 2,
		},
		{
			destination:       1122,
			hotelIDs:          []string{},
			expectHotelCounts: 1,
		},
	}

	for _, tc := range tests {
		hIDs := map[string]struct{}{}
		for _, id := range tc.hotelIDs {
			hIDs[id] = struct{}{}
		}
		data, err := paperflies.GetData(tc.destination, hIDs)
		ass.Nil(err)
		ass.Equal(tc.expectHotelCounts, len(data))
	}
}
func TestParseAmenitiesToAmenityList(t *testing.T) {
	ass := assert.New(t)
	unCodedGeneralAmenity := "UnCodedGeneralAmenity"
	unCodedRoomAmenity := "UnCodedGeneralRoomAmenity"
	amenities := paperflies.Amenities{
		General: []string{
			"outdoor pool", "indoor pool", "business center", "childcare", "parking",
			"bar", "dry cleaning", "wifi", "breakfast", "concierge", unCodedGeneralAmenity,
		},
		Room: []string{
			"tv", "coffee machine", "kettle", "hair dryer", "iron",
			"minibar", "bathtub", "aircon", unCodedRoomAmenity,
		},
	}

	expectGeneral := amen.AmenityList{}
	expectGeneral.Add(
		amen.OutdoorPool, amen.IndoorPool, amen.BusinessCenter, amen.Childcare,
		amen.Parking, amen.Bar, amen.DryCleaning, amen.Wifi, amen.Breakfast, amen.Concierge,
	)
	expectRoom := amen.AmenityList{}
	expectRoom.Add(
		amen.Tv, amen.CoffeeMachine, amen.Kettle, amen.HairDryer, amen.Iron,
		amen.Minibar, amen.Bathtub, amen.Aircon,
	)
	expectOthers := amen.AmenityList{}
	expectOthers.Add(unCodedGeneralAmenity, unCodedRoomAmenity)

	general, room, others := paperflies.ParseAmenitiesToAmenityList(amenities)
	ass.Equal(true, reflect.DeepEqual(general, expectGeneral))
	ass.Equal(true, reflect.DeepEqual(room, expectRoom))
	ass.Equal(true, reflect.DeepEqual(others, expectOthers))
}
