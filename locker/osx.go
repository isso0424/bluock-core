package locker

import (
	"os/exec"
)


type OSXLockerOperator struct {
}


func (op OSXLockerOperator) ExecuteLock() (err error) {
	err = exec.Command("pmset", "displaysleepnow").Run()
	return err
}
