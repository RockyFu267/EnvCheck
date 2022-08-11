package ssh

import (
	"log"
	"testing"
)

func Test_SSH(t *testing.T) {
	newTestCli := NewSSHClient("root", "xxxxxxx", "11.110.162.166", "22")
	res, err := newTestCli.Run("date")
	if err != nil {
		log.Println(err)
	}
	log.Println(res)

	_, err = newTestCli.UploadFile("/Users/fuao/Desktop/code/github/EnvCheck/config.yaml", "/tmp/config01.yaml")
	if err != nil {
		log.Println(err)
		return
	}
	res1, err := newTestCli.Run("chmod 777 /tmp/config01.yaml")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res1)
}
