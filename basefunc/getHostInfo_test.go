package basefunc

import (
	"fmt"
	"log"
	"testing"
)

func Test_GetHostInfo(t *testing.T) {

	res, err := GetHostInfo()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res.CPUARCH)
	fmt.Println(res.HostName)
	fmt.Println(res.OperatingSystem)

}

func Test_CheckHostName(t *testing.T) {
	str := "aa{a-a.a"
	res := CheckHostName(str)
	fmt.Println(res)
}
