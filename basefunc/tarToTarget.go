package basefunc

import (
	"EnvCheck/basecmd"
	"log"
)

func TarTarget(tgzFileList []string, sourcepath string, targetpath string) error {
	for _, v := range tgzFileList {
		_, err := basecmd.CmdAndChangeDirToResAllInOne(sourcepath, "tar -xvf "+v+" -C  "+targetpath)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
