package basefunc

import (
	"log"
	"os"
	"runtime"
)

//GetOsInfo 获取基础信息
func GetOsInfo() (OSInfo, error) {
	var tempInfo OSInfo
	tmpOs := runtime.GOOS
	tmpArch := runtime.GOARCH

	tempInfo.OperatingSystem = tmpOs
	tempInfo.CPUARCH = tmpArch

	//获取主机名
	tempHostName, err := os.Hostname()
	if err != nil {
		log.Println(err)
		//***结束***
		return tempInfo, err
	}
	tempInfo.HostName = tempHostName

	return tempInfo, nil
}
