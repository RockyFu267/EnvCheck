package controller

import (
	"fmt"
	"testing"
)

func Test_ReadConfig(t *testing.T) {
	res := ReadConfig("/Users/fuao/Desktop/code/github/EnvCheck")
	fmt.Println(res)
	fmt.Println(res.MasterIP, res.MasterPort, res.RemotePath)
}
