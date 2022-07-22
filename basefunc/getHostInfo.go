package basefunc

import (
	"log"
	"os"
	"runtime"
	"strings"
)

//GetHostInfo 获取基础信息
func GetHostInfo() (HostInfo, error) {
	var tempInfo HostInfo

	//获取架构并赋值
	tmpArch := GetCPUArch()
	tempInfo.CPUARCH = tmpArch

	//获取主机名并赋值
	tempHostName, tempHostNameCheck, err := GetHostName()
	if err != nil {
		log.Println(err)
		//***结束***
		//改进  不能结束，继续收集其他信息
		return tempInfo, err
	}
	tempInfo.HostName.HostNameStr = tempHostName
	tempInfo.HostName.CheckRes = tempHostNameCheck

	//获取是物理机还是虚拟机
	tempInfo.getHypervisor()
	if tempInfo.Hypervisor == "" {
		tempInfo.Hypervisor = "Physical"
	} else {
		tempInfo.Hypervisor = "Virtual-" + tempInfo.Hypervisor
	}
	tempInfo.getOSInfo()

	return tempInfo, nil
}

//GetHostName 获取主机名
func GetHostName() (string, bool, error) {
	res, err := os.Hostname()
	if err != nil {
		log.Println(err)
		return "", false, err
	}
	//检查主机名规范
	resCheck := CheckHostName(res)
	return res, resCheck, nil
}

//CheckHostName 检查主机名是否只包含a-z,0-9,-,.
func CheckHostName(input string) bool {
	f := func(r rune) bool {
		// return (r < 'A' && r > '9') || r > 'z' || (r > 'Z' && r < 'a') || r < '0'
		return r < '-' || (r > '.' && r < '0') || (r > '9' && r < 'a') || r > 'z'

	}
	return strings.IndexFunc(input, f) == -1

}

//GetCPUArch 获取cpu架构
func GetCPUArch() string {
	tmpArch := runtime.GOARCH
	return tmpArch
}
