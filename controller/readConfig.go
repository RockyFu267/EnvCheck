package controller

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

//ReadConfig 读取外部配置文件
func ReadConfig(input string) EnvConf {
	yamlFile, err := ioutil.ReadFile(input + "/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	var confDemo EnvConf
	err = yaml.Unmarshal(yamlFile, &confDemo)
	if err != nil {
		log.Println(err)
	}

	return confDemo

}
