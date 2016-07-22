// +build linux

package immortal

import (
	"fmt"
	"os"
	"time"
)

func (self *Daemon) watchPid(pid int, ch chan<- error) {
	if !self.isRunning(pid) {
		ch <- fmt.Errorf("PID NOT FOUND")
		return
	}

	for {
		if _, err := os.Stat(fmt.Sprintf("/proc/%d", pid)); os.IsNotExist(err) {
			ch <- fmt.Errorf("EXIT")
			return
		}
		time.Sleep(time.Second)
	}
}