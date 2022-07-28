package basefunc

import (
	bc "EnvCheck/basecmd"
	gpu "EnvCheck/gpuid"
	"log"
	"regexp"
)

// GPU information.
type GPU struct {
	Nvidia  bool     `json:"nvdia"`             //是否有nvdia显卡
	Count   int      `json:"count,omitempty"`   //数量
	Model   []string `json:"model,omitempty"`   //显卡型号
	Nouveau bool     `json:"nouveau,omitempty"` //是否安装nouveau
	Driver  bool     `json:"driver,omitempty"`  //是否安装驱动
}

//getGPUInfo 获取GPU信息
func (si *HostInfo) getGPUInfo() {
	//是否存在英伟达显卡
	resNvidia, err := bc.CmdAndChangeDirToResAllInOne("./", "lspci |grep -i nvidia | awk '{print $7}'")
	if err != nil {
		log.Println("Get Nvidia error: ", err)
		return
	}
	if len(resNvidia) == 0 {
		si.GPU.Nvidia = false
		return
	} else {
		//如果查到英伟达显卡信息
		si.GPU.Nvidia = true
	}
	//显卡数量
	si.GPU.Count = len(resNvidia)
	//根据ID获取对应的型号
	for _, v := range resNvidia {
		tmpStr := findModel(v)
		si.GPU.Model = append(si.GPU.Model, tmpStr)
	}
	//是否安装了nouveau
	resNouveau, err := bc.CmdAndChangeDirToResAllInOne("./", "lsmod | grep nouveau")
	if err != nil {
		log.Println("Get Nouveau error: ", err)
		return
	}
	if len(resNouveau) == 0 {
		si.GPU.Nouveau = false
	} else {
		si.GPU.Nouveau = true
	}
	//是否安装了驱动
	resDriver, err := bc.CmdAndChangeDirToResAllInOne("./", "nvidia-smi")
	if err != nil {
		log.Println("Get Driver error: ", err)
		si.GPU.Nouveau = false
		return
	}
	if len(resDriver) == 0 {
		si.GPU.Nouveau = false
	} else {
		si.GPU.Nouveau = true
	}

}

//findModel 根据ID找到显卡型号
func findModel(input string) string {
	filterStr := `(?m)^` + input + `.*\s`
	reg := regexp.MustCompile(filterStr)

	strList := reg.FindAllString(gpu.GPUMap, -1)
	//如果查询不到
	if len(strList) == 0 {
		log.Println("Get gpuinfo-mode error: need update ID-Map")
		return "unknown"
	}
	Model := strList[0]
	Model = DeleteExtraSpace(Model)

	//标记截取位置
	var kTmp int
	for k, v := range Model {
		if v == ' ' {
			kTmp = k
			break
		}
	}

	return Model[kTmp+1:]

}
