package controller

import (
	"io/ioutil"
	"log"

	ebf "EnvCheck/basefunc"

	"gopkg.in/yaml.v3"
)

//ReadConfig 读取外部配置文件
func ReadConfig(input string) EnvConf {
	yamlFile, err := ioutil.ReadFile(input + "/config.yaml")
	if err != nil {
		log.Println("open config file error:", err)
	}
	var confDemo EnvConf
	err = yaml.Unmarshal(yamlFile, &confDemo)
	if err != nil {
		log.Println("Unmarshal Config Error", err)
	}
	ebf.LenHostList = len(confDemo.Host)
	for k, v := range confDemo.Host {
		if v.Port == "" {
			confDemo.Host[k].Port = "22"
		}
	}
	lenTmp := len(confDemo.RemotePath)
	if string(confDemo.RemotePath[lenTmp-1]) != "/" {
		confDemo.RemotePath = confDemo.RemotePath + "/"
	}

	return confDemo

}
