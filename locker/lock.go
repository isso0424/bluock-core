package locker

import (
	"log"
	"runtime"
)

var (
	runtimeOS = runtime.GOOS
)


func Get() LockOperator {
	switch runtimeOS {
		case "linux":
			return LinuxLockerOperator{}
		case "darwin":
			return OSXLockerOperator{}
		case "windows":
			unimplementedPanic()
		default:
			unimplementedPanic()
	}

	return nil
}


// Kill myself with arguing panic when this function is called
func unimplementedPanic() {
	log.Panicf("UnimplementedError: This software does not support your os '%s'", runtimeOS)
}
