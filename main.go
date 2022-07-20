package main

import (
	bf "EnvCheck/basefunc"
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	resJson, _ := json.Marshal(res)
	fmt.Println(string(resJson))
	// fmt.Println(res)

}
