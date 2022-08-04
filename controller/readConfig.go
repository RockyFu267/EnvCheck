package controller

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func ReadConfig(input string) (res EnvConf) {
	// 读取文件所有内容装到 []byte 中
	bytes, err := ioutil.ReadFile(input + "/config.yaml")
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
	return confDemo

}
