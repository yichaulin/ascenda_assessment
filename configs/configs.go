package configs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"ascenda_assessment/logger"
)

type Configs struct {
	Suppliers              Suppliers              `yaml:"suppliers"`
	SupplierDataPriorities SupplierDataPriorities `yaml:"supplier_data_priorities"`
}

type Suppliers struct {
	ACME       string `yaml:"acme"`
	Patagonia  string `yaml:"patagonia"`
	Paperflies string `yaml:"paperflies"`
}

type SupplierDataPriorities struct {
	HotelName        map[string]uint `yaml:"hotel_name"`
	HotelAddress     map[string]uint `yaml:"hotel_address"`
	HotelDescription map[string]uint `yaml:"hotel_description"`
	HotelCountry     map[string]uint `yaml:"hotel_country"`
	HotelCity        map[string]uint `yaml:"hotel_city"`
	HotelLatLng      map[string]uint `yaml:"hotel_lat_lng"`
}

var Cfg Configs

const (
	configFileDir = "./configs"
)

func init() {
	env := "development"
	if _env := os.Getenv("ENV"); len(_env) > 0 {
		env = _env
	}

	filePath := fmt.Sprintf("%s/%s.yaml", configFileDir, env)
	fileByte, err := os.ReadFile(filePath)

	if err != nil {
		logErrorAndExit(err)
	}

	err = yaml.Unmarshal(fileByte, &Cfg)
	if err != nil {
		logErrorAndExit(err)
	}
}

func logErrorAndExit(err error) {
	logger.Error(err)
	os.Exit(1)
}
