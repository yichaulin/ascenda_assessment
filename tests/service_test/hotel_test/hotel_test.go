package hotel_test

import (
	"ascenda_assessment/configs"
	"ascenda_assessment/services/hotel"
	"ascenda_assessment/tests/mock"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetHotels(t *testing.T) {
	ass := assert.New(t)

	suppliers := configs.Cfg.Suppliers
	mockDataPaths := [][2]string{
		{suppliers.ACME, "./mockAcmeData.json"},
		{suppliers.Paperflies, "./mockPaperfliesData.json"},
		{suppliers.Patagonia, "./mockPatagoniaData.json"},
	}
	for _, p := range mockDataPaths {
		url, path := p[0], p[1]
		mockBody, err := os.ReadFile(path)
		ass.Nil(err, fmt.Sprintf("Read supplier mock data file failed. Path: %s", path))
		mock.MockAPI("GET", url, string(mockBody), http.StatusOK)
	}
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		destination       string
		hotelIDs          []string
		expectHotelCounts int
		expectError       error
	}{
		{
			destination:       "5432",
			hotelIDs:          []string{"SjyX", "iJhz", "InvalidID"},
			expectHotelCounts: 2,
			expectError:       nil,
		},
		{
			destination:       "1122",
			hotelIDs:          []string{},
			expectHotelCounts: 1,
			expectError:       nil,
		},
		{
			destination:       "",
			hotelIDs:          []string{},
			expectHotelCounts: 0,
			expectError:       hotel.InputInvalidError,
		},
	}

	for _, tc := range tests {
		hotels, err := hotel.GetHotels(tc.destination, tc.hotelIDs)
		ass.Equal(tc.expectError, err)
		ass.Equal(tc.expectHotelCounts, len(hotels))
	}
}
