package jenkinsapi

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func userAgent() string {
	sys := getSysInfo()
	return fmt.Sprintf("jenkinsapi/%s (%s/%s/%s;%s)", Version, sys.name,
		sys.release, sys.machine, runtime.Version())
}

type sysInfo struct {
	name    string // OS name such as windows/Linux
	release string // OS version 2.6.32-220.23.2.ali1089.el5.x86_64 etc
	machine string // CPU type amd64/x86_64
}

// getSysInfo gets system info
// gets the OS information and CPU type
func getSysInfo() sysInfo {
	name := runtime.GOOS
	release := "-"
	machine := runtime.GOARCH
	if out, err := exec.Command("uname", "-s").CombinedOutput(); err == nil {
		name = string(bytes.TrimSpace(out))
	}
	if out, err := exec.Command("uname", "-r").CombinedOutput(); err == nil {
		release = string(bytes.TrimSpace(out))
	}
	if out, err := exec.Command("uname", "-m").CombinedOutput(); err == nil {
		machine = string(bytes.TrimSpace(out))
	}
	return sysInfo{name: name, release: release, machine: machine}
}
