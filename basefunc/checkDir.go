package basefunc

import (
	"EnvCheck/basecmd"
	"log"
	"os"
)

func CheckDir(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		// return  err
		//如果不存在
		err = CreateDir(path)
		if err != nil {
			return err
		}
	}
	err = CreateDir(path + "/chartyaml/")
	if err != nil {
		return err
	}
	return nil
}

//CreateDir path是绝对路径
func CreateDir(dir string) error {
	_, err := basecmd.CmdAndChangeDirToResAllInOne("./", "mkdir -p "+dir)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//CreateDirGoOS path是绝对路径,不支持多层目录创建，基于go的OS包
func CreateDirGoOS(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		log.Println(err)
		return err
	}
	//如果不赋权限会有问题
	os.Chmod("path", 0755)
	return nil
}
