package configurator

import (
	"blumaton/bluock-core/configurator/inputter/mac_address"
	"fmt"
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
