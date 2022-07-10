package basecmd

import (
	"log"
	"testing"
)

func Test_CmdAndChangeDirToShow(t *testing.T) {
	params := []string{"-l"}
	err := CmdAndChangeDirToShow("/Users/fuao/Downloads/chart-test/test01", "ls", params)
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}

}
