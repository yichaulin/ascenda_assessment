package supplier_data_filter

import "ascenda_assessment/utils/string_list"

type SupplierData interface {
	GetHotelID() string
	GetDestinationID() uint64
}

func IsMatchDestinationAndHotelID(supplierData SupplierData, destination uint64, hotelIDs string_list.StringList) bool {
	hotelID := supplierData.GetHotelID()
	destinationID := supplierData.GetDestinationID()
	_, matchHotelID := hotelIDs[hotelID]

	return (matchHotelID && destinationID == destination) ||
		(len(hotelIDs) == 0 && destinationID == destination) ||
		(matchHotelID && destination == 0)
}
