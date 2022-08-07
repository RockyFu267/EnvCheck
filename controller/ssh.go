package controller

import (
	essh "EnvCheck/ssh"
	"log"
)

//SSHClient 复制client 并远程启动
func (shost *HostPara) SSHClient(pwd string, remotePath string) {
	newTestCli := essh.NewSSHClient(shost.User, shost.Password, shost.IP, shost.Port)
	//复制配置文件
	_, err := newTestCli.UploadFile(pwd+"/config.yaml", remotePath+"config.yaml")
	if err != nil {
		log.Println("copy config error: ", err)
		return
	}
	//复制工具
	_, err = newTestCli.UploadFile(pwd+"/envcheck", remotePath+"envcheck")
	if err != nil {
		log.Println("copy client error: ", err)
		return
	}
	log.Println(shost.IP + ":复制完成")

	//赋权
	cmdTmp := "chmod 777 " + remotePath + "config.yaml"
	_, err = newTestCli.Run(cmdTmp)
	if err != nil {
		log.Println("chmod config error: ", err)
		return
	}
	//赋权
	cmdTmp = "chmod 777 " + remotePath + "envcheck"
	_, err = newTestCli.Run(cmdTmp)
	if err != nil {
		log.Println("chmod client error: ", err)
		return
	}
	log.Println(shost.IP + ":赋权完成")

	//启动client
	cmdTmp = "cd " + remotePath + " && ./envcheck -role client"
	_, err = newTestCli.Run(cmdTmp)
	if err != nil {
		log.Println("start client error: ", err)
		return
	}
	log.Println(shost.IP + ":启动完成")
}
