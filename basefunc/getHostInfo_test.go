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
	fmt.Println(res.HostName)

}

func Test_CheckHostName(t *testing.T) {
	str := "aa{a-a.a"
	res := CheckHostName(str)
	fmt.Println(res)
}

func Test_getGPUInfo(t *testing.T) {
	var si HostInfo
	si.getGPUInfo()

}

func Test_DeleteExtraSpace(t *testing.T) {
	tmpstr := "     asdasd 		[      ]			asdasd  ]]]]  asdasd   asdas  "
	str := DeleteExtraSpace(tmpstr)
	fmt.Println(str)
}

func Test_afterColon(t *testing.T) {
	tmpstr := "172.0.0.1:443:"
	str := afterColon(tmpstr)
	fmt.Println(str)

}

func Test_beforeDash(t *testing.T) {
	tmpstr := "4.17.11-1.el7.elrepo.x86_64"
	str := beforeDash(tmpstr)
	fmt.Println(str)

}
