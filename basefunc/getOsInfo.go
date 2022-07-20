package basefunc

import (
	"log"
	"os"
	"runtime"
)

//GetHostInfo 获取基础信息
func GetHostInfo() (HostInfo, error) {
	var tempInfo HostInfo

	//获取架构并赋值
	tmpArch := GetCPUArch()
	tempInfo.CPUARCH = tmpArch

	//获取主机名并赋值
	tempHostName, err := GetHostName()
	if err != nil {
		log.Println(err)
		//***结束***
		//改进  不能结束，继续收集其他信息
		return tempInfo, err
	}
	tempInfo.HostName = tempHostName

	tempInfo.getHypervisor()
	if tempInfo.Hypervisor == "" {
		tempInfo.Hypervisor = "Physical"
	} else {
		tempInfo.Hypervisor = "Virtual-" + tempInfo.Hypervisor
	}

	return tempInfo, nil
}

//GetHostName 获取主机名
func GetHostName() (string, error) {
	res, err := os.Hostname()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res, nil
}

//GetCPUArch 获取cpu架构
func GetCPUArch() string {
	tmpArch := runtime.GOARCH
	return tmpArch
}
