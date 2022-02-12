package patagonia_test

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	"ascenda_assessment/apis/suppliers/patagonia"
	"ascenda_assessment/configs"
	"ascenda_assessment/tests/mock"
	amen "ascenda_assessment/utils/amenities"
)

func TestGetData(t *testing.T) {
	ass := assert.New(t)

	url := configs.Cfg.Suppliers.Patagonia
	mockDataPath := "./mockData.json"
	mockBody, err := os.ReadFile(mockDataPath)
	ass.Nil(err, fmt.Sprintf("Read Patagonia mock data file failed. Path: %s", mockDataPath))

	mock.MockAPI("GET", url, string(mockBody), http.StatusOK)
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		destination       uint64
		hotelIDs          []string
		expectHotelCounts int
	}{
		{
			destination:       5432,
			hotelIDs:          []string{"iJhz", "InvalidID"},
			expectHotelCounts: 1,
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
		data, err := patagonia.GetData(tc.destination, hIDs)
		ass.Nil(err)
		ass.Equal(tc.expectHotelCounts, len(data))
	}
}
func TestParseAmenitiesToAmenityList(t *testing.T) {
	ass := assert.New(t)
	unCodedAmenity := "UnCodedAmenity"
	amenities := []string{
		"Aircon", "Tv", "Coffee machine", "Kettle", "Hair dryer", "Iron", "Tub", "Bar", unCodedAmenity,
	}

	expectGeneral := amen.AmenityList{}
	expectGeneral.Add(amen.Bar)
	expectRoom := amen.AmenityList{}
	expectRoom.Add(amen.Aircon, amen.Tv, amen.CoffeeMachine, amen.Kettle, amen.HairDryer, amen.Iron, amen.Bathtub)
	expectOthers := amen.AmenityList{}
	expectOthers.Add(unCodedAmenity)

	general, room, others := patagonia.ParseAmenitiesToAmenityList(amenities)
	ass.Equal(true, reflect.DeepEqual(general, expectGeneral))
	ass.Equal(true, reflect.DeepEqual(room, expectRoom))
	ass.Equal(true, reflect.DeepEqual(others, expectOthers))
}
