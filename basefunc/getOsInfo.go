package basefunc

import (
	"runtime"
)

//GetOsInfo 获取基础信息
func GetOsInfo() OSInfo {
	var tempInfo OSInfo
	tmpOs := runtime.GOOS
	tmpArch := runtime.GOARCH

	tempInfo.OperatingSystem = tmpOs
	tempInfo.CPUARCH = tmpArch

	return tempInfo
}
