package utils

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

//Windows, Linux
const (
	Windows = `windows`
	Linux   = `linux`
	Other   = `other`
)

//VerifyOS ...
func VerifyOS() string {
	sysType := runtime.GOOS
	switch sysType {
	case `windows`:
		return Windows
	case `linux`:
		return Linux
	default:
		return Other
	}
}

//GetCPUAndMemoryInfo ...
func GetCPUAndMemoryInfo() {
	cpuInfo, _ := cpu.Percent(time.Second, false)
	memInfo, _ := mem.VirtualMemory()
	fmt.Println(cpuInfo[0], memInfo.UsedPercent)
}
