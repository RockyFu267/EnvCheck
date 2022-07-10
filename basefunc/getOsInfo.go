package basefunc

import "runtime"

//GetOsInfo 获取基础信息
func GetOsInfo() OSInfo {
	var tempInfo OSInfo
	os := runtime.GOOS
	arch := runtime.GOARCH

	tempInfo.OS = os
	tempInfo.ARCH = arch
	return tempInfo
}
