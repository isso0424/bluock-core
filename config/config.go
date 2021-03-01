package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RSSIThreshold int  `yaml:"rssi_threshold"`
	MACAddresses []string `yaml:"mac_addresses"`
}

func loadFromFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func LoadConfig(configPath string) (c Config, err error) {
	data, err := loadFromFile(configPath)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &c)

	return
}
