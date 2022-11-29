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
	str := findModel("0049")
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
func Test_resIOPS(t *testing.T) {
	tmpstr := "write: IOPS=3211, BW=50.2MiB/s (52.6MB/s)(1004MiB/20003msec); 0 zone resets"
	str1, str2 := resIOPS(tmpstr)
	fmt.Println(str1)
	fmt.Println(str2)

	tmpstr1 := "read: IOPS=3087, BW=48.2MiB/s (50.6MB/s)(965MiB/20003msec)"
	str3, str4 := resIOPS(tmpstr1)
	fmt.Println(str3)
	fmt.Println(str4)

}

func Test_notFoundLib64(t *testing.T) {
	var tmpstr = []string{`	linux-vdso.so.1 (0x00007ffe3b9de000)`, `	libtcmalloc_minimal.so.4 => not found`, `	libnuma.so.1 => /lib64/libnuma.so.1 (0x00007f4ef063d000)`, `	librt.so.1 => /lib64/librt.so.1 (0x00007f4ef0434000)`, `	libz.so.1 => /lib64/libz.so.1 (0x00007f4ef021d000)`, `	libpthread.so.0 => /lib64/libpthread.so.0 (0x00007f4eefffd000)`, `	libm.so.6 => /lib64/libm.so.6 (0x00007f4eefc7b000)`, `	libdl.so.2 => /lib64/libdl.so.2 (0x00007f4eefa77000)`, `	libibverbs.so.1 => not found`, `	librdmacm.so.1 => not found`, `	libc.so.6 => /lib64/libc.so.6 (0x00007f4eef6b5000)`, `	/lib64/ld-linux-x86-64.so.2 (0x00007f4ef0849000)`}

	str := notFoundLib64(tmpstr)
	fmt.Println(str)

}
