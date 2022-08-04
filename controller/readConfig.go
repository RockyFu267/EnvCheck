package controller

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func readConfig() {
	// 读取文件所有内容装到 []byte 中
	bytes, err := ioutil.ReadFile("/Users/fuao/Desktop/code/github/EnvCheck/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	// 创建配置文件的结构体
	var confDemo EnvConf
	// 调用 Unmarshall 去解码文件内容
	// 注意要穿配置结构体的指针进去
	err = yaml.Unmarshal(bytes, &confDemo)
	if err != nil {
		log.Println(err)
	}
	for k, v := range confDemo.Host {
		fmt.Println(k, v)
	}

}
