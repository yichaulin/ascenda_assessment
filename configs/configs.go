package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"

	"ascenda_assessment/logger"
)

type Configs struct {
	Suppliers              map[string]Supplier
	SupplierDataPriorities *SupplierDataPriorities `yaml:"supplier_data_priorities"`
}

type Supplier struct {
	Name    string `yaml:"name"`
	Url     string `yaml:"url"`
	Enabled bool   `yaml:"enabled"`
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

func init() {
	_, absFilePath, _, _ := runtime.Caller(0)
	configFileDir := filepath.Dir(absFilePath)

	env := "development"
	if _env := os.Getenv("ENV"); len(_env) > 0 {
		env = _env
	}

	filePath := fmt.Sprintf("%s/%s.yaml", configFileDir, env)
	fileByte, err := os.ReadFile(filePath)
	if err != nil {
		logErrorAndExit(err)
	}

	suppliers := struct {
		suppliers []Supplier `yaml:"suppliers"`
	}{}
	err = yaml.Unmarshal(fileByte, &suppliers)
	for _, s := range suppliers.suppliers {
		Cfg.Suppliers[s.Name] = s
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
