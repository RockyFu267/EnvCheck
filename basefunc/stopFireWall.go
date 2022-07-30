package basefunc

import (
	bc "EnvCheck/basecmd"
	"log"
)

//stopFireWall 停止防火墙
func (si *HostInfo) stopFireWall() {
	_, err := bc.CmdAndChangeDirToResAllInOne("./", "systemctl stop firewalld.service")
	if err != nil {
		log.Println("Get StopFirewalld error: ", err)
		si.Others.Firewall.Stop = false
	}
	_, err = bc.CmdAndChangeDirToResAllInOne("./", "systemctl disable firewalld.service")
	if err != nil {
		log.Println("Get disableFirewalld error: ", err)
		si.Others.Firewall.Disable = false
		return
	}
	si.Others.Firewall.Stop = true
	si.Others.Firewall.Disable = true
}
