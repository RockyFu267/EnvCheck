package basefunc

import (
	"log"
	"testing"
)

func Test_CheckDir(t *testing.T) {
	err := CheckDir("/Users/fuao/Downloads/chart-test/test01/002/0002/00002")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
}

func Test_CreateDir(t *testing.T) {
	err := CreateDir("/etc/fuao/Downloads/chart-test/test01/001/0001/000001/")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
}

func Test_CreateDirGoOS(t *testing.T) {
	err := CreateDirGoOS("/Users/fuao/Downloads/chart-test/test01/002/0002/000002/")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
}
