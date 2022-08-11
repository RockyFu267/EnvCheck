package controller

import (
	ebf "EnvCheck/basefunc"
	"log"
)

//SSHClient 复制client 并远程启动
func (shost *HostPara) SSHClient(pwd string, remotePath string, mastetIp string, masterPort string) {
	var tmpPostInfo ebf.PostInfo
	tmpPostInfo.PostIP = shost.IP
	newTestCli := ebf.NewSSHClient(shost.User, shost.Password, shost.IP, shost.Port)
	//测试与master端口联通性
	cmdTmp := "curl http://" + mastetIp + ":" + masterPort + "/health_check/"
	curlResTmp, err := newTestCli.Run(cmdTmp)
	if err != nil {
		log.Println("copy config error: ", err)
		tmpPostInfo.Res = false
		ebf.PostInfoList = append(ebf.PostInfoList, tmpPostInfo)
		return
	}
	//debug
	log.Println(curlResTmp, "11111111111111")
	//复制配置文件
	_, err = newTestCli.UploadFile(pwd+"/config.yaml", remotePath+"config.yaml")
	if err != nil {
		log.Println("copy config error: ", err)
		tmpPostInfo.Res = false
		ebf.PostInfoList = append(ebf.PostInfoList, tmpPostInfo)
		return
	}
	//复制工具
	_, err = newTestCli.UploadFile(pwd+"/envcheck", remotePath+"envcheck")
	if err != nil {
		log.Println("copy client error: ", err)
		tmpPostInfo.Res = false
		ebf.PostInfoList = append(ebf.PostInfoList, tmpPostInfo)
		return
	}
	log.Println(shost.IP + ":复制完成")

	//赋权
	cmdTmp = "chmod 777 " + remotePath + "config.yaml"
	_, err = newTestCli.Run(cmdTmp)
	if err != nil {
		log.Println("chmod config error: ", err)
		tmpPostInfo.Res = false
		ebf.PostInfoList = append(ebf.PostInfoList, tmpPostInfo)
		return
	}
	//赋权
	cmdTmp = "chmod 777 " + remotePath + "envcheck"
	_, err = newTestCli.Run(cmdTmp)
	if err != nil {
		log.Println("chmod client error: ", err)
		tmpPostInfo.Res = false
		ebf.PostInfoList = append(ebf.PostInfoList, tmpPostInfo)
		return
	}
	log.Println(shost.IP + ":赋权完成")

	//启动client
	cmdTmp = "cd " + remotePath + " && ./envcheck -role client"
	_, err = newTestCli.Run(cmdTmp)
	if err != nil {
		log.Println("start client error: ", err)
		tmpPostInfo.Res = false
		ebf.PostInfoList = append(ebf.PostInfoList, tmpPostInfo)
		return
	}
	log.Println(shost.IP + ":启动完成")

}
