package controller

import (
	essh "EnvCheck/ssh"
	"log"
)

//SSHClient 复制client 并远程启动
func (shost *HostPara) SSHClient(pwd string, remotePath string) {
	newTestCli := essh.NewSSHClient(shost.User, shost.Password, shost.IP, shost.Port)
	//复制配置文件
	_, err := newTestCli.UploadFile(pwd+"/config.yaml", remotePath)
	if err != nil {
		log.Println(err)
		return
	}
	//复制工具
	_, err = newTestCli.UploadFile(pwd+"/envcheck", remotePath)
	if err != nil {
		log.Println(err)
		return
	}

	//赋权
	cmdTmp := "chmod 777 " + remotePath + "config.yaml"
	res, err := newTestCli.Run(cmdTmp)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res)
	//赋权
	cmdTmp = "chmod 777 " + remotePath + "envcheck"
	res, err = newTestCli.Run(cmdTmp)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res)

	//启动client
	cmdTmp = "cd " + remotePath + " && nohup ./envcheck &"
	_, err = newTestCli.Run(cmdTmp)
	if err != nil {
		log.Println(err)
		return
	}
}
