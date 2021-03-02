package locker

import "os/exec"


type LinuxLockerOperator struct {
}


func (op LinuxLockerOperator) ExecuteLock() (err error) {
	err = exec.Command("loginctl", "lock-session").Run()

	return
}
