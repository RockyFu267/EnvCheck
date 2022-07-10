package basefunc

import (
	"EnvCheck/basecmd"
	"log"
	"os"
)

func ReadDirList(targetpath string) ([]string, error) {
	resTmp, err := basecmd.CmdAndChangeDirToResAllInOne(targetpath, "ls -l|grep  '^d' | awk '(NR > 1){print $9}'")
	if err != nil {
		log.Println(err)
		return resTmp, err
	}
	res := []string{}
	//过滤有效目录
	for _, v := range resTmp {
		_, err := os.Stat(targetpath + "/" + v + "/Chart.yaml")
		if err != nil {
			log.Println("get os.stat fail: ", err)
			//如果不存在
			continue
		}
		res = append(res, v)
	}
	return res, nil
}
