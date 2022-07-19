package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	pwdPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
		//***结束***
		return
	}
	fmt.Println("当前路径：", pwdPath)
	abc := runtime.GOARCH
	abc1 := runtime.GOOS
	fmt.Println(abc, abc1)
}
