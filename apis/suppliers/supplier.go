package supplier

import (
	"fmt"

	"ascenda_assessment/apis/suppliers/acme"
	"ascenda_assessment/apis/suppliers/paperflies"
	"ascenda_assessment/apis/suppliers/patagonia"
)

func GetData(supplier string, destination uint64, hotelIDs map[string]struct{}) (interface{}, error) {
	switch supplier {
	case acme.SupplierName:
		return acme.GetData(destination, hotelIDs)
	case patagonia.SupplierName:
		return patagonia.GetData(destination, hotelIDs)
	case paperflies.SupplierName:
		return paperflies.GetData(destination, hotelIDs)
	default:
		return nil, fmt.Errorf("Invalid supplier: %s", supplier)
	}
}

func GetAllSupplierNames() []string {
	return []string{
		acme.SupplierName, patagonia.SupplierName, paperflies.SupplierName,
	}
}
