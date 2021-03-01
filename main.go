package main

import (
	"blumaton/bluock-core/config"
	"blumaton/bluock-core/scanner"
)

func main() {
	c, err := config.LoadConfig("config/example.yml")
	if err != nil {
		panic(err)
	}

	err = scanner.ScanDevices(&c)
	if err != nil {
		panic(err)
	}
}
