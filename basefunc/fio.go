package basefunc

import (
	bc "EnvCheck/basecmd"
	"log"
)

//

//yumFio yum安装fio
func yumFio() (res bool) {
	_, err := bc.CmdAndChangeDirToResAllInOne("./", "yum install fio -y")
	if err != nil {
		log.Println(": ", err)
		return false
	}
	log.Println("fio already installed")
	return true
}

//cmdFioCheck 检查fio是否已安装
func cmdFioCheck() bool {
	resCmdfioCheck, err := bc.CmdAndChangeDirToResAllInOne("./", "fio --version")
	if err != nil {
		log.Println("fio not installed: ", err)
		return false
	}
	log.Println("fio --version: ", resCmdfioCheck)
	return true
}
