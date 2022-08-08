package controller

import (
	ebf "EnvCheck/basefunc"
)

func WriteChan() {

}

//CheckInfoList 检查是否已全员检查完毕
func CheckInfoList() bool {
	//比较已执行长度与目标长度
	lenPostHost := len(ebf.PostInfoList)
	return lenPostHost >= ebf.LenHostList
}
