package scanner

import (
	"blumaton/bluock-core/config"
	"log"
	"os"
	"os/signal"
	"syscall"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func handleFinish(sc chan os.Signal) {
	<-sc
	adapter.StopScan()
}

func ScanDevices(c *config.Config) (err error) {
	defer log.Println("Stop scanning")
	err = adapter.Enable()
	if err != nil {
		return
	}

	log.Println("Start device scanning")
	log.Printf("Handling addresses: %v", c.MACAddresses)
	log.Printf("RSSI threshold: %d", c.RSSIThreshold)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	go handleFinish(sc)

	err = adapter.Scan(func (adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		// TODO: Create handler and use here.
	})

	return
}
