package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	pwdPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
		//***结束***
		return
	}
	fmt.Println("当前路径：", pwdPath)

}
