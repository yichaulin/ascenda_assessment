package hotel_test

import (
	"ascenda_assessment/apis/client"
	"ascenda_assessment/apis/suppliers/acme"
	"ascenda_assessment/apis/suppliers/paperflies"
	"ascenda_assessment/apis/suppliers/patagonia"
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
	mockSupplierData(ass)
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

func TestGetHotelsWithAddressPriorityConfig(t *testing.T) {
	ass := assert.New(t)
	mockSupplierData(ass)
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		priority        []uint // [{acme}, {patagonia}, {paperflies}, ...]
		expectedAddress string
	}{
		{
			priority:        []uint{1, 2, 3},
			expectedAddress: "8 Sentosa Gateway, Beach Villas, 098269",
		},
		{
			priority:        []uint{4, 2, 3},
			expectedAddress: "8 Sentosa Gateway, Beach Villas",
		},
	}

	for _, tc := range tests {
		supplierDataPriorities := configs.Cfg.SupplierDataPriorities
		addressPriority := map[string]uint{
			acme.SupplierName:       tc.priority[0],
			patagonia.SupplierName:  tc.priority[1],
			paperflies.SupplierName: tc.priority[2],
		}
		supplierDataPriorities.HotelAddress = addressPriority

		hotels, err := hotel.GetHotels("5432", []string{"iJhz"})
		ass.Nil(err)
		ass.Equal(1, len(hotels))
		ass.Equal(*hotels[0].Location.Address, tc.expectedAddress)
	}

}

func mockSupplierData(ass *assert.Assertions) {
	suppliers := configs.Cfg.Suppliers
	mockConfig := []struct {
		url          string
		mockDataPath string
		apiClient    client.Client
	}{
		{
			url:          suppliers.ACME,
			mockDataPath: "./mockAcmeData.json",
			apiClient:    acme.ApiClient,
		},
		{
			url:          suppliers.Paperflies,
			mockDataPath: "./mockPaperfliesData.json",
			apiClient:    paperflies.ApiClient,
		},
		{
			url:          suppliers.Patagonia,
			mockDataPath: "./mockPatagoniaData.json",
			apiClient:    patagonia.ApiClient,
		},
	}

	for _, c := range mockConfig {
		url, path, client := c.url, c.mockDataPath, c.apiClient
		mockBody, err := os.ReadFile(path)
		ass.Nil(err, fmt.Sprintf("Read supplier mock data file failed. Path: %s", path))
		mock.MockAPI(client, "GET", url, string(mockBody), http.StatusOK)
	}
}
