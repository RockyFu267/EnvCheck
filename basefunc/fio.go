package basefunc

import (
	bc "EnvCheck/basecmd"
	"log"
	"strings"
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

//DiskIOTest 执行fio验证并测试
func (si *HostInfo) DiskIOTest(input string) {
	//检查fio是否安装
	if cmdFioCheck() {
		log.Println("fio installed")
		_, err := si.CmdFioTest(input)
		if err != nil {
			log.Println(err)
			return
		}
		return
	} else {
		//fio安装
		if yumFio() {
			log.Println("fio installed")
			_, err := si.CmdFioTest(input)
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

//CmdFioTest 执行fio命令测试获取结果
func (si *HostInfo) CmdFioTest(input string) (FioResInfo, error) {
	pathTmp := input
	randWrite := "fio -filename=" + pathTmp + " -direct=1 -iodepth 1 -thread -rw=randwrite -ioengine=psync -bs=16k -size=1G -numjobs=10 -runtime=20 -group_reporting -name=mytest"
	randRead := "fio -filename=" + pathTmp + " -direct=1 -iodepth 1 -rw=randread -ioengine=psync -bs=16k -size=1G -numjobs=10 -runtime=20 -group_reporting -name=mytest"
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
	for _, v := range resFioRadnWrite {
		arrayTMP := strings.Fields(DeleteExtraSpace(v))
		if len(arrayTMP) < 2 {
			continue
		}
		if arrayTMP[0] == "write:" {
			si.Storage.DiskIO.RandWrite.IOPS, si.Storage.DiskIO.RandWrite.BW = resIOPS(DeleteExtraSpace(v))
			break
		}
	}
	for _, v := range resFioRadnRead {
		arrayTMP := strings.Fields(DeleteExtraSpace(v))
		if len(arrayTMP) < 2 {
			continue
		}
		if arrayTMP[0] == "read:" {
			si.Storage.DiskIO.RandRead.IOPS, si.Storage.DiskIO.RandRead.BW = resIOPS(DeleteExtraSpace(v))
			break
		}
	}

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

//resIOPS 过滤随机写的结果
func resIOPS(input string) (iops string, bw string) {
	// if BeforeColon(input) == write
	arrayTMP := strings.Fields(input)
	for _, v := range arrayTMP {
		if len(v) <= 6 {
			continue
		}
		log.Println(v)
		log.Println(v[0:4])
		if v[0:5] == "IOPS=" {
			iops = v[5 : len(v)-1]
		}
		if v[0:3] == "BW=" {
			bw = v[3:]
		}
	}
	return iops, bw
}
