package supplier

import (
	"ascenda_assessment/apis/suppliers/acme"
	"ascenda_assessment/apis/suppliers/paperflies"
	"ascenda_assessment/apis/suppliers/patagonia"
)

func GetData(supplier string, destination uint64, hotelIDs map[string]struct{}) (data interface{}, err error) {
	switch supplier {
	case acme.SupplierName:
		data, err = acme.GetData(destination, hotelIDs)
	case patagonia.SupplierName:
		data, err = patagonia.GetData(destination, hotelIDs)
	case paperflies.SupplierName:
		data, err = paperflies.GetData(destination, hotelIDs)
	}

	return data, err
}

func GetAllSupplierNames() []string {
	return []string{
		acme.SupplierName, patagonia.SupplierName, paperflies.SupplierName,
	}
}
