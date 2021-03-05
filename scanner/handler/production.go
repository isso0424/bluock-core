package handler

import (
	"blumaton/bluock-core/config"
	"log"

	"tinygo.org/x/bluetooth"
)

type productionHandler struct {
	c *config.Config
}

func New(c *config.Config) productionHandler {
	return productionHandler{c}
}

func (handler productionHandler) lockScreenValidator(addr string) bool {
	for _, targetAddr := range handler.c.MACAddresses {
		if addr == targetAddr {
			return true
		}
	}

	return false
}

func (handler productionHandler) OnDetected(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
	if handler.c.RSSIThreshold > int(device.RSSI) && handler.lockScreenValidator(device.Address.String()) {
		err := handler.c.Locker.ExecuteLock()
		if err != nil {
			log.Println("[error]", err)
		}
	}
}
