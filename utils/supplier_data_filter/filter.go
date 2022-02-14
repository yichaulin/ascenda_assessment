package supplier_data_filter

type SupplierData interface {
	GetHotelID() string
	GetDestinationID() uint64
}

func IsMatchDestinationAndHotelID(supplierData SupplierData, destination uint64, hotelIDs map[string]struct{}) bool {
	hotelID := supplierData.GetHotelID()
	destinationID := supplierData.GetDestinationID()
	_, matchHotelID := hotelIDs[hotelID]

	return (matchHotelID && destinationID == destination) ||
		(len(hotelIDs) == 0 && destinationID == destination) ||
		(matchHotelID && destination == 0)
}
