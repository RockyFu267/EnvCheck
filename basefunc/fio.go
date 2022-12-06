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
		if installFio() {
			log.Println("fio installed")
			_, err := si.CmdFioTestLocal(input)
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

//CmdFioTestLocal 执行fio命令测试获取结果
func (si *HostInfo) CmdFioTestLocal(input string) (FioResInfo, error) {
	pathTmp := input
	randWrite := "./fio-lib64/fio -filename=" + pathTmp + " -direct=1 -iodepth 1 -thread -rw=randwrite -ioengine=psync -bs=16k -size=1G -numjobs=10 -runtime=20 -group_reporting -name=mytest"
	randRead := "./fio-lib64/fio -filename=" + pathTmp + " -direct=1 -iodepth 1 -rw=randread -ioengine=psync -bs=16k -size=1G -numjobs=10 -runtime=20 -group_reporting -name=mytest"
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

//installFio 安装fio
func installFio() (res bool) {
	//检查目录是否存在
	checkDir, err := CheckDir("fio-lib64")
	if err != nil {
		log.Println("Check fio-lib64 ERROR: ", err)
		//***结束***
		return false
	}
	//如果不存在就解压
	if !checkDir {
		resTarLib64 := cmdFioTarLib64()
		if !resTarLib64 {
			return false
		}
	}
	//使用ldd获取结果
	resLddFio, err := lddFio()
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println(resLddFio)
	//分析结果缺少的动态库 数组
	listNotFound := notFoundLib64(resLddFio)
	//如果不缺 那就直接可用并验证
	if len(listNotFound) == 0 {
		if cmdFioCheckLocal() {
			log.Println("fio installed")
		} else {
			log.Println("fio install-Error")
			return false
		}
	} else {
		//缺少的动态库 复制到 默认的动态库路径
		for _, v := range listNotFound {
			err := cpFioLib64(v)
			if err != nil {
				log.Println(err)
				return false
			}
		}
		if cmdFioCheckLocal() {
			log.Println("fio installed")
		} else {
			log.Println("fio install-Error")
			return false
		}

	}

	return true
}

//notFoundLib64 获取结果缺少的动态库
func notFoundLib64(inputdata []string) []string {
	var res []string
	for _, v := range inputdata {
		tmpStr := DeleteExtraSpace(v)
		if len(tmpStr) <= 13 {
			continue
		}
		if tmpStr[len(tmpStr)-9:] == "not found" {
			res = append(res, tmpStr[:len(tmpStr)-13])
		}
	}
	return res
}

//cpFioLib64 动态库 复制到 默认的动态库路径
func cpFioLib64(input string) (err error) {
	cmd := "cp ./fio-lib64/" + input + " /usr/lib64/"
	_, err = bc.CmdAndChangeDirToResAllInOne("./", cmd)
	if err != nil {
		log.Println("cp "+input+" error: ", err)
		return err
	}
	log.Println("cp " + input + " success")
	return nil
}

//lddFio 动态库检查
func lddFio() (res []string, err error) {
	res, err = bc.CmdAndChangeDirToResAllInOne("./", "./fio-lib64/ldd ./fio-lib64/fio")
	if err != nil {
		log.Println("ldd error: ", err)
		return res, err
	}
	log.Println("ldd success")
	return res, nil

}

// //yumFio yum安装fio  暂不使用这种方式
// func yumFio() (res bool) {
// 	_, err := bc.CmdAndChangeDirToResAllInOne("./", "yum install fio -y")
// 	if err != nil {
// 		log.Println(": ", err)
// 		return false
// 	}
// 	log.Println("fio already installed")
// 	return true
// }

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

//cmdFioCheckLocal 检查fio是否已安装
func cmdFioCheckLocal() bool {
	resCmdfioCheck, err := bc.CmdAndChangeDirToResAllInOne("./", "./fio-lib64/fio --version")
	if err != nil {
		log.Println("fio not installed: ", err)
		return false
	}
	log.Println("fio --version: ", resCmdfioCheck)
	return true
}

//cmdFioTarLib64 解压tar包
func cmdFioTarLib64() bool {
	resCmdfioTar, err := bc.CmdAndChangeDirToResAllInOne("./", "tar -zxvf fio-lib64.tar")
	if err != nil {
		log.Println("lib64 tar error: ", err)
		return false
	}
	log.Println("tar: ", resCmdfioTar)
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
