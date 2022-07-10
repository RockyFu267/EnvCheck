package basefunc

import (
	"fmt"
	"log"
	"testing"
)

func Test_ReadFileListNew(t *testing.T) {
	res, err := ReadFileListNew("/Users/fuao/Downloads/chart-test/test01")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println(res)

}
func Test_ReadTGZFileList(t *testing.T) {
	res, err := ReadTGZFileList("/Users/fuao/Downloads/chart-test/test01")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println(res)

}
func Test_ReadFileList(t *testing.T) {
	res, err := ReadFileList("/Users/fuao/Downloads/chart-test/test01")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println(res)

}

func Test_TestList(t *testing.T) {
	res := TestList()

	fmt.Println(res)

}
