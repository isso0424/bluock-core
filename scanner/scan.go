package scanner

import (
	"blumaton/bluock-core/scanner/handler"
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

func ScanDevices(scanningHandler handler.BluetoothHandler) (err error) {
	defer log.Println("Stop scanning")
	err = adapter.Enable()
	if err != nil {
		return
	}

	log.Println("Start device scanning")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	go handleFinish(sc)

	err = adapter.Scan(scanningHandler.OnDetected)

	return
}
