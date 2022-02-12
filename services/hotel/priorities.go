package hotel

import (
	"ascenda_assessment/apis/suppliers/acme"
	"ascenda_assessment/apis/suppliers/paperflies"
	"ascenda_assessment/apis/suppliers/patagonia"
)

var (
	namePriorityConfig = map[string]uint{
		acme.SupplierName:       1,
		patagonia.SupplierName:  2,
		paperflies.SupplierName: 3,
	}

	addressPriorityConfig = map[string]uint{
		acme.SupplierName:       1,
		patagonia.SupplierName:  2,
		paperflies.SupplierName: 3,
	}

	descriptionPriorityConfig = map[string]uint{
		acme.SupplierName:       1,
		patagonia.SupplierName:  2,
		paperflies.SupplierName: 3,
	}

	countryPriorityConfig = map[string]uint{
		acme.SupplierName:       1,
		patagonia.SupplierName:  2,
		paperflies.SupplierName: 3,
	}

	cityPriorityConfig = map[string]uint{
		acme.SupplierName:       1,
		patagonia.SupplierName:  2,
		paperflies.SupplierName: 3,
	}

	latLngPriorityConfig = map[string]uint{
		acme.SupplierName:       1,
		patagonia.SupplierName:  2,
		paperflies.SupplierName: 3,
	}
)
