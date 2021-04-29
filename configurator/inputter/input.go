package configurator

import (
	"blumaton/bluock-core/configurator/inputter/mac_address"
	"blumaton/bluock-core/configurator/inputter/rssi"
	"fmt"
	"strconv"
)

func inputText() (string, error) {
	var address string
	_, err := fmt.Scan(address)
	if err != nil {
		return "", err
	}

	return address, nil
}

func InputMacAddress() (mac_address.MacAddress, error) {
	fmt.Printf("Please input MAC address\n>>> ")
	text, err := inputText()
	if err != nil {
		return "", err
	}

	return mac_address.New(text)
}

func InputRSSI() (rssi.RSSI, error) {
	fmt.Printf("Please input RSSI threshold\n>>> ")
	text, err := inputText()
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}

	return rssi.New(num)
}
