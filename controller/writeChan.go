package controller

import (
	ebf "EnvCheck/basefunc"
	"fmt"
	"log"
)

func WriteChan() {

}

//CheckInfoList 检查是否已全员检查完毕
func CheckInfoList() bool {
	fmt.Println("")
	log.Println("")
	lenConfigHost := len(ebf.HostInfoList)
	lenPostHost := len(ebf.PostInfoList)
	return lenPostHost >= lenConfigHost
}
