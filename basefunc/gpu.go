package basefunc

import (
	gpu "EnvCheck/gpuid"
	"fmt"
	"regexp"
)

//getGPUInfo 获取GPU信息
func getGPUInfo() {
	fmt.Println(gpu.GPUMap)
	reg := regexp.MustCompile(`(?m)^25aa.*\s`)

	fmt.Println(reg.FindAllString(gpu.GPUMap, -1))
}
