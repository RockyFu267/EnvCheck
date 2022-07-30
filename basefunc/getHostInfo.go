package basefunc

import (
	bc "EnvCheck/basecmd"
	"log"
	"os"
	_ "runtime" //暂时不用
	"strings"
)

//GetHostInfo 获取基础信息
func GetHostInfo() (HostInfo, error) {
	var tempInfo HostInfo
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
	//获取内存相关信息
	tempInfo.getMemoryInfo()
	tempInfo.Memory.Size = tempInfo.Memory.Size / 1024
	//获取显卡相关信息
	tempInfo.getGPUInfo()
	//获取网卡相关信息
	tempInfo.getNetworkInfo()
	//获取磁盘信息
	tempInfo.getStorageInfo()
	//获取系统时区
	tempInfo.GetTimeZone()
	//关闭防火墙
	tempInfo.stopFireWall()
	//检查软件以及端口
	tempInfo.serviceCheck()
	//获取版本信息
	tempInfo.getMetaInfo()

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

//GetTimeZone 获取主机名
func (si *HostInfo) GetTimeZone() {
	res, err := bc.CmdAndChangeDirToResAllInOne("./", "date +%Z")
	if err != nil {
		log.Println(err)
		return
	}
	if len(res) == 0 {
		return
	}
	si.Others.TimeZone = res[0]

}

//CheckHostName 检查主机名是否只包含a-z,0-9,-,.
func CheckHostName(input string) bool {
	f := func(r rune) bool {
		// return (r < 'A' && r > '9') || r > 'z' || (r > 'Z' && r < 'a') || r < '0'
		return r < '-' || (r > '.' && r < '0') || (r > '9' && r < 'a') || r > 'z'

	}
	return strings.IndexFunc(input, f) == -1

}

//get 获取磁盘分区
//df -Th | grep -v devtmpfs | grep -v tmpfs |  grep -v overlay | grep data | awk 'NR>1'      只关注data
