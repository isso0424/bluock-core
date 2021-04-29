package configurator

import (
	"blumaton/bluock-core/config"
	"blumaton/bluock-core/configurator/inputter"
	"blumaton/bluock-core/configurator/writer"
)

func UpdateConfig() error {
	macAddr, err := inputter.InputMacAddress()
	if err != nil {
		return err
	}

	rssi, err := inputter.InputRSSI()
	if err != nil {
		return err
	}

	c := config.Config{
		RSSIThreshold: int(rssi),
		MACAddresses: []string{string(macAddr)},
	}

	return writer.WriteConfig(c)
}
