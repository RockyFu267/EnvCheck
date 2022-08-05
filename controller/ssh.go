package controller

import (
	bc "EnvCheck/basecmd"
	"log"
)

//SCPClient 复制client 到client角色列表  这一段是屎山待自己实现ssh、scp后再修改
func (shost *HostPara) SCPClient() {
	if shost.Port == "" {
		shost.Port = "22"
	}
	scpCmdClient := "/root/envcheck/gossh" + " -t push -h " + shost.IP + " -P " + shost.Port + " -u root -p " + shost.Password + " -f /root/envcheck/envcheck /root/ "
	err := SCPFile(scpCmdClient)
	if err != nil {
		log.Println(err)
		return
	}

	scpCmdConfig := "/root/envcheck/gossh" + " -t push -h " + shost.IP + " -P " + shost.Port + " -u root -p " + shost.Password + " -f /root/envcheck/config.yaml /root/ "
	err = SCPFile(scpCmdConfig)
	if err != nil {
		log.Println(err)
		return
	}
	scpCmdClientSh := "/root/envcheck/gossh" + " -t push -h " + shost.IP + " -P " + shost.Port + " -u root -p " + shost.Password + " -f /root/envcheck/client-start.sh /root/ "
	err = SCPFile(scpCmdClientSh)
	if err != nil {
		log.Println(err)
		return
	}
}

//SshStartClient 远程启动client
func (shost *HostPara) SshStartClient() {
	if shost.Port == "" {
		shost.Port = "22"
	}
	startCmdClient := "/root/envcheck/gossh" + " -t cmd -h " + shost.IP + " -P " + shost.Port + " -u root -p " + shost.Password + " -f /root/client-start.sh "
	resStart, err := bc.CmdAndChangeDirToResAllInOne("./", startCmdClient)
	if err != nil {
		log.Println("start "+shost.IP+" error: ", err)
		return
	}
	if len(resStart) < 3 {
		log.Println("start "+shost.IP+" error:", err)
		return
	}
	if resStart[2] != "return=0" {
		log.Println("start "+shost.IP+" error:", err)
		return
	}

}

func SCPFile(cpCmd string) error {
	resSCP, err := bc.CmdAndChangeDirToResAllInOne("./", cpCmd)
	if err != nil {
		log.Println("cp file error: ", err)
		return err
	}
	if len(resSCP) < 3 {
		log.Println("cp file error: ", err)
		return err
	}
	if resSCP[2] != "return=0" {
		log.Println("cp file error: ", err)
		return err
	}
	return nil
}
