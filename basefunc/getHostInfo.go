package basefunc

import (
	bc "EnvCheck/basecmd"
	"log"
	"os"
	"runtime"
	"strconv"
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

	//获取CPU相关信息
	tempInfo.getCPUInfo()
	//获取内核相关信息
	tempInfo.getKernelInfo()

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

//CheckDir 检查文件或者目录是否存在
func CheckDir(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

//getMemorySie 获取内存大小
func getMemorySie() (int64, error) {
	resSize, err := bc.CmdAndChangeDirToResAllInOne("./", "cat /Users/fuao/Downloads/proc/meminfo | grep 'MemTotal' | awk '{print $2}'")
	if err != nil {
		log.Println("Get MemorySize Error: ", err)
		// return 0, err
		return 0, err
	}
	resStr := resSize[0]
	intTemp, err := strconv.ParseInt(resStr, 10, 64)
	if err != nil {
		log.Println("Get MemorySize Error: ", err)
		// return 0, err
		return 0, err
	}
	resInt := intTemp / 1024
	if resInt < 1024 {
		return 1, nil
	}
	resInt = resInt / 1024
	return resInt, nil
}
