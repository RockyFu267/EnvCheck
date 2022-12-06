package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func Test_ReadConfig(t *testing.T) {
	res, err := ReadConfig("/Users/fuao/Desktop/code/github/EnvCheck")
	if err != nil {
		log.Println(err)
	}
	resJson, _ := json.MarshalIndent(res, "", " ")
	fmt.Println(string(resJson))
	fmt.Println(res.Host[0].DiskCheck)

}

func Test_ReadResJson(t *testing.T) {
	res, err := ReadResJson("/Users/fuao/Downloads/2022-09-04_17-47-47.json")
	fmt.Println(err)
	resJson, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		if err != nil {
			log.Println("解析数组错误", err)
			return
		}
	}
	fmt.Println(string(resJson))
}
