package handler

import "tinygo.org/x/bluetooth"

type BluetoothHandler interface {
	OnDetected(adapter *bluetooth.Adapter, device bluetooth.ScanResult)
}
