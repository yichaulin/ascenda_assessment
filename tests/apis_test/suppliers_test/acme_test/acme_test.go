package acme_test

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	"ascenda_assessment/apis/suppliers/acme"
	"ascenda_assessment/configs"
	"ascenda_assessment/tests/mock"
	amen "ascenda_assessment/utils/amenities"
)

func TestGetData(t *testing.T) {
	ass := assert.New(t)

	url := configs.Cfg.Suppliers.ACME
	mockDataPath := "./mockData.json"
	mockBody, err := os.ReadFile(mockDataPath)
	ass.Nil(err, fmt.Sprintf("Read ACME mock data file failed. Path: %s", mockDataPath))

	mock.MockAPI(acme.ApiClient, "GET", url, string(mockBody), http.StatusOK)
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
		data, err := acme.GetData(tc.destination, hIDs)
		ass.Nil(err)
		ass.Equal(tc.expectHotelCounts, len(data))
	}
}

func TestParseFacilitiesToAmenityList(t *testing.T) {
	ass := assert.New(t)
	unCodedAmenity := "UnCodedAmenity"
	facilities := []string{
		"Pool", "BusinessCenter", "DryCleaning",
		"Breakfast", "Bar", "Aircon", "BathTub", "WiFi",
		unCodedAmenity,
	}

	expectGeneral := amen.AmenityList{}
	expectGeneral.Add(amen.Pool, amen.BusinessCenter, amen.DryCleaning, amen.Breakfast, amen.Bar, amen.Wifi)
	expectRoom := amen.AmenityList{}
	expectRoom.Add(amen.Aircon, amen.Bathtub)
	expectOthers := amen.AmenityList{}
	expectOthers.Add(unCodedAmenity)

	general, room, others := acme.ParseFacilitiesToAmenityList(facilities)
	ass.Equal(true, reflect.DeepEqual(general, expectGeneral))
	ass.Equal(true, reflect.DeepEqual(room, expectRoom))
	ass.Equal(true, reflect.DeepEqual(others, expectOthers))
}
