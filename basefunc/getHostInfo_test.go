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
	str := "gzbh-intelmbx005.gzbh.baidu.com"
	res := CheckHostName(str)
	fmt.Println(res)
}

func Test_getGPUInfo(t *testing.T) {
	var si HostInfo
	si.getGPUInfo()

}

func Test_findModel(t *testing.T) {
	str := findModel("1bb3")
	fmt.Println(str)
}

func Test_findModelID(t *testing.T) {
	str := findModelID(`   af:00.0 3D controller [0302]: NVIDIA Corporation Device [10de:1bb3] (rev a1)  `)
	fmt.Println(str)
}

func Test_DeleteExtraSpace(t *testing.T) {
	tmpstr := "     asdasd 		[      ]			asdasd  ]]]]  asdasd   asdas  "
	str := DeleteExtraSpace(tmpstr)
	fmt.Println(str)
}

func Test_afterColon(t *testing.T) {
	tmpstr := "172.0.0.1443:asd"
	str := afterColon(tmpstr)
	fmt.Println(str)

}

func Test_BeforeDash(t *testing.T) {
	tmpstr := "4.17.11-1.el7.elrepo.x86_64"
	str := BeforeDash(tmpstr)
	fmt.Println(str)

}

func Test_BeforeColon(t *testing.T) {
	tmpstr := ":asdasd"
	str := BeforeColon(tmpstr)
	fmt.Println(str)
	fmt.Println(len(str))

}
