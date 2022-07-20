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
