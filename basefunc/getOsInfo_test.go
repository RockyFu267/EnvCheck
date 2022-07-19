package basefunc

import (
	"fmt"
	"log"
	"testing"
)

func Test_GetOsInfo(t *testing.T) {

	res, err := GetOsInfo()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res.CPUARCH)
	fmt.Println(res.HostName)
	fmt.Println(res.OperatingSystem)

}
