package basefunc

import (
	gpu "EnvCheck/gpuid"
	"fmt"
	"regexp"
)

// GPU information.
type GPU struct {
	Nvidia  bool   `json:"nvdia"`           //是否是nvdia
	Model   string `json:"model,omitempty"` //显卡型号
	Nouveau bool   `json:"nouveau"`         //是否安装nouveau
	Driver  bool   `json:"driver"`          //是否安装驱动
}

//getGPUInfo 获取GPU信息
func (si *HostInfo) getGPUInfo() error {
	reg := regexp.MustCompile(`(?m)^1bb3.*\s`)

	strList := reg.FindAllString(gpu.GPUMap, -1)
	Model := strList[0]
	fmt.Println(Model)
	Model = DeleteExtraSpace(Model)
	fmt.Println(Model)
	return nil
}
