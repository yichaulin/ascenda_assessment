package acme

import (
	"encoding/json"
	"strings"
	"time"

	"ascenda_assessment/apis/client"
	"ascenda_assessment/configs"

	amen "ascenda_assessment/utils/amenities"
	"ascenda_assessment/utils/string_list"
	"ascenda_assessment/utils/supplier_data_filter"
)

type ACMEData struct {
	ID            string   `json:"Id"`
	DestinationID uint64   `json:"DestinationId"`
	Name          string   `json:"Name"`
	Latitude      *float32 `json:"Latitude"`
	Longitude     *float32 `json:"Longitude"`
	Address       *string  `json:"Address"`
	City          *string  `json:"City"`
	Country       *string  `json:"Country"`
	PostalCode    *string  `json:"PostalCode"`
	Description   *string  `json:"Description"`
	Facilities    []string `json:"Facilities"`
}

const (
	SupplierName = "acme"

	Pool           = "Pool"
	BusinessCenter = "BusinessCenter"
	DryCleaning    = "DryCleaning"
	Breakfast      = "Breakfast"
	Bar            = "Bar"
	WiFi           = "WiFi"

	Aircon  = "Aircon"
	BathTub = "BathTub"
)

var ApiClient client.Client

func init() {
	ApiClient = client.New()
	if cfg, ok := configs.Cfg.Suppliers[SupplierName]; ok && cfg.Timeout > 0 {
		ApiClient.SetTimeout(cfg.Timeout * time.Second)
	}
}

func (a ACMEData) GetHotelID() string {
	return a.ID
}

func (a ACMEData) GetDestinationID() uint64 {
	return a.DestinationID
}

func GetData(destination uint64, hotelIDs string_list.StringList) (interface{}, error) {
	acmeData := []ACMEData{}
	supplierConfig, ok := configs.Cfg.Suppliers[SupplierName]
	if !ok || !supplierConfig.Enabled {
		return acmeData, nil
	}

	url := configs.Cfg.Suppliers[SupplierName].Url
	resp, err := ApiClient.Get(url)
	if err != nil {
		return acmeData, err
	}

	tmp := []ACMEData{}
	json.Unmarshal(resp.Body(), &tmp)

	acmeData = make([]ACMEData, 0, len(tmp))
	for _, hotel := range tmp {
		if supplier_data_filter.IsMatchDestinationAndHotelID(hotel, destination, hotelIDs) {
			acmeData = append(acmeData, hotel)
		}
	}

	return acmeData, nil
}

func ParseFacilitiesToAmenityList(facilities []string) (general string_list.StringList, room string_list.StringList, others string_list.StringList) {
	general = string_list.StringList{}
	room = string_list.StringList{}
	others = string_list.StringList{}

	for _, f := range facilities {
		f := strings.TrimSpace(f)
		switch f {
		case Pool:
			general.Add(amen.Pool)
		case Breakfast:
			general.Add(amen.Breakfast)
		case BusinessCenter:
			general.Add(amen.BusinessCenter)
		case WiFi:
			general.Add(amen.Wifi)
		case DryCleaning:
			general.Add(amen.DryCleaning)
		case Aircon:
			room.Add(amen.Aircon)
		case BathTub:
			room.Add(amen.Bathtub)
		case Bar:
			general.Add(amen.Bar)
		default:
			others.Add(f)
		}
	}

	return general, room, others
}
