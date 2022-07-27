package basefunc

import (
	"os"
	"regexp"
	"strings"
)

//DeleteExtraSpace 删除多余空格 删除行首行位空格 删除换行符 替换[]
func DeleteExtraSpace(s string) string {
	//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "  ", " ", -1) //替换tab为空格
	//删除行首行位空格 删除换行符 替换[]
	s1 = strings.Replace(s1, "[", " ", -1)
	s1 = strings.Replace(s1, "]", " ", -1)
	s1 = strings.Replace(s1, "\n", "", -1)
	s1 = strings.TrimSpace(s1)
	regstr := "\\s{2,}"                          //两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)             //编译正则表达式
	s2 := make([]byte, len(s1))                  //定义字符数组切片
	copy(s2, s1)                                 //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 {                     //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}

	return string(s2)
}

//CheckDir 检查文件或者目录是否存在
func CheckDir(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// //getMemorySie 获取内存大小 手动检查文件
// func getMemorySie() (int64, error) {
// 	resSize, err := bc.CmdAndChangeDirToResAllInOne("./", "cat /Users/fuao/Downloads/proc/meminfo | grep 'MemTotal' | awk '{print $2}'")
// 	if err != nil {
// 		log.Println("Get MemorySize Error: ", err)
// 		// return 0, err
// 		return 0, err
// 	}
// 	resStr := resSize[0]
// 	intTemp, err := strconv.ParseInt(resStr, 10, 64)
// 	if err != nil {
// 		log.Println("Get MemorySize Error: ", err)
// 		// return 0, err
// 		return 0, err
// 	}
// 	resInt := intTemp / 1024
// 	if resInt < 1024 {
// 		return 1, nil
// 	}
// 	resInt = resInt / 1024
// 	return resInt, nil
// }
