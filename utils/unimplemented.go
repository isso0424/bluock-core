package utils

import (
	"log"
	"runtime"
)

var (
	runtimeOS = runtime.GOOS
)


// Kill myself with arguing panic when this function is called
func UnimplementedPanic() {
	log.Panicf("UnimplementedError: This software does not support your os '%s'", runtimeOS)
}
