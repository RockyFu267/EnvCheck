package main

import (
	bf "EnvCheck/basefunc"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	pwdPath, err := os.Getwd()
	if err != nil {
		log.Println("Get pwdPATH ERROR: ", err)
		//***结束***
		return
	}
	fmt.Println("当前路径：", pwdPath)
	//获取本机信息
	res, err := bf.GetHostInfo()
	if err != nil {
		log.Println("Get OsInfo ERROR: ", err)
		//***结束***
		return
	}
	// resJson, _ := json.Marshal(res)
	resJson, _ := json.MarshalIndent(res, "", " ")
	fmt.Println(string(resJson))

	//获取当前时间
	timeStr := time.Now().Format("2006-01-02_15:04:05")
	envFileName := "./" + timeStr + ".json"

	//文件是否存在
	checkDir, err := bf.CheckDir(envFileName)
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
			fmt.Println("创建失败", err)
			return
		}
	} else {
		//不存在 创建文件
		file, err = os.Create(envFileName)
		if err != nil {
			fmt.Println("创建失败", err)
			return
		}
	}
	defer file.Close()
	//写入文件
	_, err = io.WriteString(file, string(resJson))
	if err != nil {
		fmt.Println("写入错误：", err)
		return
	}
	fmt.Println("写入成功")
}
