package controller

import (
	ebf "EnvCheck/basefunc"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//WriteRes 写入结果
func WriteRes() {
	//获取当前时间
	timeStr := time.Now().Format("2006-01-02_15-04-05")
	envFileName := "./" + timeStr + ".json"

	//文件是否存在
	checkDir, err := ebf.CheckDir(envFileName)
	if err != nil {
		log.Println("Get jsonfile ERROR: ", err)
		//***结束***
		return
	}
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())
	tmpStr := strconv.Itoa(rand.Intn(10000))
	var file *os.File
	if checkDir {
		//创建文件 名字后添加随机数
		envFileNameTmp := "./" + timeStr + "_" + tmpStr + ".json"
		file, err = os.Create(envFileNameTmp)
		if err != nil {
			log.Println("创建失败", err)
			return
		}
	} else {
		//不存在 创建文件
		file, err = os.Create(envFileName)
		if err != nil {
			log.Println("创建失败", err)
			return
		}
	}
	defer file.Close()
	resJson, err := json.MarshalIndent(ebf.HostInfoList, "", " ")
	if err != nil {
		if err != nil {
			log.Println("解析数组错误", err)
			return
		}
	}
	//写入文件
	_, err = io.WriteString(file, string(resJson))
	if err != nil {
		log.Println("写入错误：", err)
		return
	}
	log.Println("写入成功")

}

//WriteResImage 写入简易结果，方便拍照获取
func WriteResImage() {
	//获取当前时间
	timeStr := time.Now().Format("2006-01-02_15-04-05")
	envFileName := "./" + timeStr + ".txt"

	//文件是否存在
	checkDir, err := ebf.CheckDir(envFileName)
	if err != nil {
		log.Println("Get jsonfile ERROR: ", err)
		//***结束***
		return
	}
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())
	tmpStr := strconv.Itoa(rand.Intn(10000))
	var file *os.File
	if checkDir {
		//创建文件 名字后添加随机数
		envFileNameTmp := "./" + timeStr + "_" + tmpStr + ".txt"
		file, err = os.Create(envFileNameTmp)
		if err != nil {
			log.Println("创建失败", err)
			return
		}
	} else {
		//不存在 创建文件
		file, err = os.Create(envFileName)
		if err != nil {
			log.Println("创建失败", err)
			return
		}
	}
	defer file.Close()
	var strRes string
	//循环获取结果
	//存放机器类型
	hypervisorMap := make(map[string]int)
	kernelMap := make(map[string]bool)
	osMap := make(map[string]bool)
	GPUModeMap := make(map[string]bool)
	GPUCount := 0
	cpuarchMap := make(map[string]bool)
	cpuCount := 0
	instructionMap := make(map[string]bool)
	diskTypeMap := make(map[string]bool)
	var momorySizeTotal uint
	netCardSpeedMap := make(map[string]bool)
	netCardSMap := make(map[string]bool)
	illegalHostMap := make(map[string]bool)
	diskResMap := make(map[string]ebf.FioResInfo)
	for _, v := range ebf.HostInfoList {
		//是否虚拟机
		// if _, ok := hypervisorMap[v.Hypervisor]; ok {
		// } else {
		// 	hypervisorMap[v.Hypervisor] = hypervisorMap[v.Hypervisor] + 1
		// }
		if v.Hypervisor == "Physical" {
			hypervisorMap["Physical"] = hypervisorMap["Physical"] + 1
		} else {
			hypervisorMap["Virtual"] = hypervisorMap["Virtual"] + 1
		}

		//获取操作系统版本
		osInfo := v.OS.Vendor + "-" + v.OS.Release
		if _, ok := osMap[osInfo]; ok {
		} else {
			osMap[osInfo] = true
		}
		//获取内核版本
		kernelStr := ebf.BeforeDash(v.Kernel.Release)
		if _, ok := kernelMap[kernelStr]; ok {
		} else {
			kernelMap[kernelStr] = true
		}
		//GPU加速卡
		GPUCount = GPUCount + v.GPU.Count
		for _, k := range v.GPU.Model {
			if _, ok := GPUModeMap[k]; ok {
			} else {
				GPUModeMap[k] = true
			}
		}
		//CPU架构
		if _, ok := cpuarchMap[v.CPU.CPUARCH]; ok {
		} else {
			cpuarchMap[v.CPU.CPUARCH] = true
		}
		cpuCount = cpuCount + int(v.CPU.Threads)
		//CPU指令集
		if v.CPU.Avx {
			instructionMap["avx"] = true
		}
		if v.CPU.Avx2 {
			instructionMap["avx2"] = true
		}
		if v.CPU.Bmi2 {
			instructionMap["bmi2"] = true
		}
		//磁盘信息
		for _, k := range v.Storage.DiskInfo {
			if _, ok := diskTypeMap[k.Type]; ok {
			} else {
				diskTypeMap[k.Type] = true
			}
		}
		//内存大小
		momorySizeTotal = momorySizeTotal + v.Memory.Size
		//网卡信息
		for _, k := range v.Network {
			if _, ok := netCardSpeedMap[k.Speed]; ok {
			} else {
				netCardSpeedMap[k.Speed] = true
			}
		}
		if len(v.Network) >= 2 {
			netCardSMap[v.Meta.IP] = true
		}
		//检查主机名是否合法
		if !v.HostName.CheckRes {
			illegalHostMap[v.Meta.IP] = true
		}

		if v.Storage.DiskIO.RandRead.BW != "" && v.Storage.DiskIO.RandRead.IOPS != "" && v.Storage.DiskIO.RandWrite.BW != "" && v.Storage.DiskIO.RandWrite.IOPS != "" {
			var diskResTmp ebf.FioResInfo
			diskResTmp.RandRead.BW = v.Storage.DiskIO.RandRead.BW
			diskResTmp.RandRead.IOPS = v.Storage.DiskIO.RandRead.IOPS
			diskResTmp.RandWrite.BW = v.Storage.DiskIO.RandWrite.BW
			diskResTmp.RandWrite.IOPS = v.Storage.DiskIO.RandWrite.IOPS
			diskResMap[v.HostName.HostNameStr] = diskResTmp
		}

	}
	//拼接字符串类型的结果输出
	//机器类型
	strPhysicalCount := strconv.Itoa(hypervisorMap["Physical"])
	strVirtualCount := strconv.Itoa(hypervisorMap["Virtual"])
	if len(hypervisorMap) >= 2 {
		strRes = "机器类型统计: 有物理机" + strPhysicalCount + "台, 有虚拟机 " + strVirtualCount + " 台" + "\n"
	} else {
		_, ok := hypervisorMap["Physical"]
		if len(hypervisorMap) == 1 && ok {
			strRes = "机器类型统计: 全部是物理机,一共 " + strPhysicalCount + " 台\n"
		} else {
			strRes = "机器类型统计: 全部是虚拟机,一共 " + strVirtualCount + " 台\n"
		}
	}
	//操作系统版本列表
	var osStr string
	for k := range osMap {
		osStr = osStr + " " + k
	}
	strRes = strRes + "操作系统版本: " + osStr + "\n"
	//内核版本列表
	var kernelStr string
	for k := range kernelMap {
		kernelStr = kernelStr + " " + k
	}
	strRes = strRes + "内核版本: " + kernelStr + "\n"
	//GPU加速卡信息
	var gpuModeStr string
	if GPUCount < 1 {
	} else {
		strGPUCount := strconv.Itoa(GPUCount)
		strRes = strRes + "GPU加速卡数量: " + strGPUCount + "\n"
		for k := range GPUModeMap {
			gpuModeStr = gpuModeStr + " " + k
		}
		strRes = strRes + "GPU加速卡型号: " + gpuModeStr + "\n"
	}
	//CPU信息
	var cpuarchStr string
	var instructionStr string
	for k := range cpuarchMap {
		cpuarchStr = cpuarchStr + " " + k
	}
	if len(cpuarchMap) >= 2 {
		strRes = strRes + "CPU架构列表: " + cpuarchStr + "\n"
	} else {
		strRes = strRes + "CPU架构: " + cpuarchStr + "\n"
	}
	strcpuCount := strconv.Itoa(cpuCount)
	strRes = strRes + "CPU总核数: " + strcpuCount + "\n"
	for k := range instructionMap {
		instructionStr = instructionStr + " " + k
	}
	strRes = strRes + "CPU指令集列表: " + instructionStr + "\n"
	//磁盘信息
	var diskTypeStr string
	for k := range diskTypeMap {
		diskTypeStr = diskTypeStr + " " + k
	}
	strRes = strRes + "磁盘类型: " + diskTypeStr + "\n"
	//内存大小
	strMemorySizeTotal := strconv.Itoa(int(momorySizeTotal))
	strRes = strRes + "内存总大小: " + strMemorySizeTotal + "\n"
	//网卡信息
	var netCardSpeedStr string
	var netCardsStr string
	for k := range netCardSpeedMap {
		netCardSpeedStr = netCardSpeedStr + " " + k
	}
	strRes = strRes + "网卡速率种类列表: " + netCardSpeedStr + "\n"
	if len(netCardSMap) <= 1 {
	} else {
		for k := range netCardSMap {
			netCardsStr = netCardsStr + " " + k
		}
		strRes = strRes + "有多张网卡的机器列表: " + netCardsStr + "\n"
	}
	//主机名小写的机器列表
	var illegalHostStr string
	if len(illegalHostMap) <= 0 {
	} else {
		for k := range illegalHostMap {
			illegalHostStr = illegalHostStr + " " + k
		}
		strRes = strRes + "需要修改的主机名的机器列表: " + illegalHostStr + "\n"
	}

	//主机名小写的机器列表
	var diskResStr string
	if len(diskResMap) > 0 {
		for k, v := range diskResMap {
			diskResStr = diskResStr + " 主机名：" + k + " 随机读带宽速率：" + v.RandRead.BW + " 随机读IOPS: " + v.RandRead.IOPS + " 随机写带宽速率：" + v.RandWrite.BW + " 随机写IOPS: " + v.RandWrite.IOPS + "\n"
		}
		strRes = strRes + "磁盘测试结果: \n" + diskResStr + "\n"
	}
	fmt.Println(strRes)

	//写入文件
	_, err = io.WriteString(file, strRes)
	if err != nil {
		log.Println("写入错误：", err)
		return
	}
	log.Println("写入成功")

}
