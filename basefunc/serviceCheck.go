package basefunc

import (
	bc "EnvCheck/basecmd"
	"fmt"
	"log"
)

type ServiceCheck struct {
	DockerInstalled bool     `json:"docker_installed"`
	K8sInstalled    bool     `json:"k8s_installed"`
	PortList        []string `json:"port_list,omitempty"`
}

//serviceCheck 检查docker、k8s 以及获取监听端口
func (si *HostInfo) serviceCheck() {
	resCheckDocker, err := CheckDir("/var/run/docker.sock")
	if err != nil {
		log.Println("Check docker error: ", err)
	} else {
		si.Others.ServiceCheck.DockerInstalled = resCheckDocker
	}

	resCheckK8s01, err := CheckDir("/var/lib/kubelet/config.yaml")
	if err != nil {
		log.Println("Check docker error: ", err)
	} else {
		si.Others.ServiceCheck.K8sInstalled = resCheckK8s01
	}

	portMap := make(map[string]bool)
	var portList []string
	resSSList, err := bc.CmdAndChangeDirToResAllInOne("./", "ss -lntp |awk 'NR>1'  | awk '{print $4}'")
	if err != nil {
		log.Println("Get port list error: ", err)
		return
	}
	for _, v := range resSSList {
		tmpV := afterColon(v)
		if _, ok := portMap[tmpV]; ok {
			continue
		} else {
			portList = append(portList, tmpV)
		}
		//
		fmt.Println(portMap, tmpV)
	}
	si.Others.ServiceCheck.PortList = portList
}
