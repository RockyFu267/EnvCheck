package basefunc

import (
	"EnvCheck/basecmd"
	"fmt"
	"log"
	"os/exec"
)

//ReadFileListNew 获取所有文件列表 可以过滤文件夹和格式
func ReadFileListNew(dir string) ([]string, error) {
	resTmp, err := basecmd.CmdAndChangeDirToResAllInOne(dir, "ls -l|grep -v '^d' | awk '(NR > 1){print $9}'")
	if err != nil {
		log.Println(err)
		return resTmp, err
	}
	res := []string{}
	//过滤格式
	for _, v := range resTmp {
		fmt.Println(v, len(v))
		if len(v) <= 4 {
			continue
		}
		if v[len(v)-4:] != ".tgz" {
			continue
		}
		res = append(res, v)
	}
	return res, nil
}

func ReadTGZFileList(dir string) ([]string, error) {
	res, err := ReadFileList(dir)
	if err != nil {
		log.Println(err)
		return res, err
	}
	for _, v := range res {
		fmt.Println(len(v))
	}
	return res, nil
}

//ReadFileList 获取所有文件列表 不能过滤文件夹
func ReadFileList(dir string) ([]string, error) {
	params := []string{}
	res, err := basecmd.CmdAndChangeDirToRes(dir, "ls", params)
	if err != nil {
		log.Println(err)
		return res, err
	}
	return res, nil
}

func TestList() string {
	cmd := "ls -l |grep -v ^d"
	cmd01 := exec.Command("bash", "-c", cmd)
	cmd01.Dir = "/Users/fuao/Downloads/chart-test/test01/test01"
	out, err := cmd01.Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd)
	}
	return string(out)
}
