package config

import (
	"blumaton/bluock-core/locker"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RSSIThreshold int                 `yaml:"rssi_threshold"`
	MACAddresses  []string            `yaml:"mac_addresses"`
	Locker        locker.LockOperator `yaml:"-"`
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

	c.Locker = locker.Get()

	return
}
