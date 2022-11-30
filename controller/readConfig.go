package controller

import (
	ebf "EnvCheck/basefunc"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

//ReadConfig 读取外部配置文件
func ReadConfig(input string) (EnvConf, error) {
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
		log.Println(k, confDemo.Disk.Type)
		switch confDemo.Disk.Type {
		case "none":
			confDemo.Host[k].DiskCheck = "disable"
		case "all":
			if v.DiskCheck == "disable" {
				continue
			}
			confDemo.Host[k].DiskCheck = "enable"
			if v.DiskCheckPath == "" {
				confDemo.Host[k].DiskCheckPath = confDemo.Disk.DiskCheckPath
			}
		case "custom":
			if v.DiskCheck == "enable" && v.DiskCheckPath == "" {
				return confDemo, errors.New("disk.type Error,please check parameter")
			}
			if v.DiskCheck == "" {
				confDemo.Host[k].DiskCheck = "disable"
			}
		default:
			return confDemo, errors.New("disk.type Error,please check parameter")
		}
	}
	lenTmp := len(confDemo.RemotePath)
	if string(confDemo.RemotePath[lenTmp-1]) != "/" {
		confDemo.RemotePath = confDemo.RemotePath + "/"
	}

	return confDemo, nil

}

//ReadResJson 读取json结果
func ReadResJson(input string) ([]ebf.HostInfo, error) {
	var res []ebf.HostInfo
	jsonFile, err := ioutil.ReadFile(input)
	if err != nil {
		log.Println("open ResJson file error:", err)
		return res, err
	}
	err = json.Unmarshal(jsonFile, &res)
	if err != nil {
		log.Println("Unmarshal ResJson Error", err)
		return res, err
	}
	return res, nil
}
