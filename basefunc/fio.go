package basefunc

import (
	bc "EnvCheck/basecmd"
	"log"
)

type FioResInfo struct {
	RandWrite FioWriteRes `json:"randwrite"`
	RandRead  FioReadRes  `json:"randread"`
}
type FioWriteRes struct {
	IOPS string `json:"iops"`
	BW   string `json:"bw"`
}

type FioReadRes struct {
	IOPS string `json:"iops"`
	BW   string `json:"bw"`
}

func diskIOTest(input string) {
	//检查fio是否安装
	if cmdFioCheck() {
		log.Println("fio installed")
	} else {
		//fio安装
		if yumFio() {
			log.Println("fio installed")
			_, err := cmdFioTest(input)
			if err != nil {
				log.Println(err)
				return
			}
		} else {
			log.Println("fio installed fail")
			//得采取其他测试方式
			return
		}
	}

}

//cmdFioTest 执行fio命令测试获取结果
func cmdFioTest(input string) (FioResInfo, error) {
	pathTmp := input
	randWrite := "fio -filename=" + pathTmp + " -direct=1 -iodepth 1 -thread -rw=randwrite -ioengine=psync -bs=16k -size=1G -numjobs=10 -runtime=10 -group_reporting -name=mytest"
	randRead := "fio -filename=" + pathTmp + " -direct=1 -iodepth 1 -rw=randread -ioengine=psync -bs=16k -size=1G -numjobs=10 -runtime=10 -group_reporting -name=mytest"
	var res FioResInfo
	resFioRadnWrite, err := bc.CmdAndChangeDirToResAllInOne("./", randWrite)
	if err != nil {
		log.Println("fio randwrite: ", err)
		return res, err
	}

	resFioRadnRead, err := bc.CmdAndChangeDirToResAllInOne("./", randRead)
	if err != nil {
		log.Println("fio randrand: ", err)
		return res, err
	}

	log.Println(resFioRadnWrite)
	log.Println(resFioRadnRead)

	return res, nil
}

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
