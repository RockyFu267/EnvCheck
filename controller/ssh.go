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
	//测试与master端口联通性 先curl 请求 写入页面信息 再校验是否与预期相符
	//远程执行目前还未找到合适的直接返回网络联通性结果的命令，等找到再优化替换
	cmdTmp := "curl http://" + mastetIp + ":" + masterPort + "/health_check/ > " + remotePath + "curl_res.txt"
	_, err := newTestCli.Run(cmdTmp)
	if err != nil {
		log.Println("curl master port error: ", err)
		tmpPostInfo.Res = false
		ebf.PostInfoList = append(ebf.PostInfoList, tmpPostInfo)
		return
	}
	cmdTmp = "cat " + remotePath + "curl_res.txt"
	curlResTmp, err := newTestCli.Run(cmdTmp)
	if err != nil {
		log.Println("get curl res error01: ", err)
		tmpPostInfo.Res = false
		ebf.PostInfoList = append(ebf.PostInfoList, tmpPostInfo)
		return
	}
	//去掉换行符
	RES := ebf.DeleteExtraSpace(curlResTmp)
	//检查获取内容是否与健康检查页内容一致
	if RES != "HelloWorld" {
		log.Println("get curl res error02: ", err)
		tmpPostInfo.Res = false
		ebf.PostInfoList = append(ebf.PostInfoList, tmpPostInfo)
		return
	}
	log.Println(shost.IP + ":网络测试通过")
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
	//复制lib64
	_, err = newTestCli.UploadFile(pwd+"/fio-lib64.tar", remotePath+"fio-lib64.tar")
	if err != nil {
		log.Println("copy fio-lib64.tar error: ", err)
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

	//判断是否要开启磁盘检测
	if shost.DiskCheck == "enable" {
		cmdTmp = "cd " + remotePath + " && ./envcheck -role client " + "-disktest " + shost.DiskCheckPath
	} else {
		cmdTmp = "cd " + remotePath + " && ./envcheck -role client"
	}
	//启动client
	// cmdTmp = "cd " + remotePath + " && ./envcheck -role client"
	_, err = newTestCli.Run(cmdTmp)
	if err != nil {
		log.Println("start client error: ", err)
		tmpPostInfo.Res = false
		ebf.PostInfoList = append(ebf.PostInfoList, tmpPostInfo)
		return
	}
	log.Println(shost.IP + ":启动完成")

}
