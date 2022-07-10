package basefunc

import (
	"log"
	"testing"
)

func Test_TarTarget(t *testing.T) {
	res, err := ReadFileListNew("/Users/fuao/Downloads/chart-test/test01")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	err = TarTarget(res, "/Users/fuao/Downloads/chart-test/test01/", "/Users/fuao/Downloads/chart-test/test01/test000")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
}
