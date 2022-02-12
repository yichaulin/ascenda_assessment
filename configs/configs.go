package configs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"ascenda_assessment/logger"
)

type Configs struct {
	Suppliers Suppliers `yml:"suppliers"`
}

type Suppliers struct {
	ACME       string `yaml:"acme"`
	Patagonia  string `yaml:"patagonia"`
	Paperflies string `yaml:"paperflies"`
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
