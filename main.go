package main

import (
	"blumaton/bluock-core/config"
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
)

type mode int

type options struct {
	mode
}

func parseArgs() options {
	args := options{mode: production}
	isConnectionTest := flag.Bool("connection-test", false, "Start bluetooth connection test")

	flag.Parse()
	if *isConnectionTest {
		args.mode = connectionTest
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
