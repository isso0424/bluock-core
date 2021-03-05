package locker

import (
	"blumaton/bluock-core/utils"
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
		utils.UnimplementedPanic()
	default:
		utils.UnimplementedPanic()
	}

	return nil
}
