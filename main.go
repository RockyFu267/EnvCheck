package main

import (
	bf "EnvCheck/basefunc"
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
	res, err := bf.GetOsInfo()
	if err != nil {
		log.Println("Get OsInfo ERROR: ", err)
		//***结束***
		return
	}
	fmt.Println(res)

}
