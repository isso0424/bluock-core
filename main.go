package main

import (
	"blumaton/bluock-core/config"
	"blumaton/bluock-core/configurator"
	"blumaton/bluock-core/scanner"
	"blumaton/bluock-core/scanner/handler"
	"blumaton/bluock-core/utils"
	"flag"
	"runtime"
)

const (
	runtimeOS  = runtime.GOOS
	production = iota
	connectionTest
	configurate
)

type mode int

type options struct {
	mode
}

func parseArgs() options {
	args := options{mode: production}
	isConnectionTest := flag.Bool("connection-test", false, "Start bluetooth connection test")
	isConfigurate := flag.Bool("configurate", false, "")

	flag.Parse()
	if *isConnectionTest {
		args.mode = connectionTest
	}

	if *isConfigurate {
		args.mode = configurate
	}

	return args
}

func main() {
	isSupportedOS()

	args := parseArgs()

	var scanningHandler handler.BluetoothHandler

	switch args.mode {
	case connectionTest:
		scanningHandler = handler.NewTester()
	case production:
		c, err := config.LoadConfig("config/example.yml")
		if err != nil {
			panic(err)
		}
		scanningHandler = handler.New(&c)
	case configurate:
		err := configurator.UpdateConfig()
		if err != nil {
			panic(err)
		}

		return
	}

	err := scanner.ScanDevices(scanningHandler)
	if err != nil {
		panic(err)
	}
}

func isSupportedOS() {
	switch runtimeOS {
	case "linux":
		return
	case "darwin":
		return
	default:
		utils.UnimplementedPanic()
	}
}
