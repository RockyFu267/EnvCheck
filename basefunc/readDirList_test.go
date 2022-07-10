package basefunc

import (
	"fmt"
	"log"
	"testing"
)

func Test_ReadDirList(t *testing.T) {
	res, err := ReadDirList("/Users/fuao/Desktop/code/github/SamllToolForDeploy/output/")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println(res)
}
