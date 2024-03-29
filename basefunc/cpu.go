package basefunc

import (
	bc "EnvCheck/basecmd"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// CPU information.
type CPU struct {
	CPUARCH         string          `json:"cpuarch,omitempty"`
	Threads uint   `json:"threads,omitempty"` // number of logical (HT) CPU cores
	Vendor  string `json:"vendor,omitempty"`
	Model   string `json:"model,omitempty"`
	Speed   uint   `json:"speed,omitempty"`   // CPU clock rate in MHz
	Cache   uint   `json:"cache,omitempty"`   // CPU cache size in KB
	Cpus    uint   `json:"cpus,omitempty"`    // number of physical CPUs
	Cores   uint   `json:"cores,omitempty"`   // number of physical CPU cores
	Avx2    bool   `json:"avx2,omitempty"`
	Avx     bool   `json:"avx,omitempty"`
	Bmi2    bool   `json:"bmi2,omitempty"`
}

var (
	reTwoColumns = regexp.MustCompile("\t+: ")
	reExtraSpace = regexp.MustCompile(" +")
	reCacheSize  = regexp.MustCompile(`^(\d+) KB$`)
)

func (si *HostInfo) getCPUInfo() {
	//获取架构并赋值   
	si.CPU.CPUARCH  = GetCPUArch()

	//arm不检查，鉴权通过容器化部署
	if si.OS.Architecture == "amd64" {
		resAvx2, err := bc.CmdAndChangeDirToResAllInOne("./", "cat /proc/cpuinfo |grep  'avx2'")
		if err != nil {
			log.Println("Get cpuinfo-ep error: ", err)
		}
		resAvx, err := bc.CmdAndChangeDirToResAllInOne("./", "cat /proc/cpuinfo |grep  'avx '")
		if err != nil {
			log.Println("Get cpuinfo-ep error: ", err)
		}
		resBmi2, err := bc.CmdAndChangeDirToResAllInOne("./", "cat /proc/cpuinfo |grep 'bmi2'")
		if err != nil {
			log.Println("Get cpuinfo-ep error: ", err)
		}
		//判断指令集的结果
		if len(resAvx2) != 0 {
			si.CPU.Avx2 = true
		}
		if len(resAvx) != 0 {
			si.CPU.Avx = true
		}
		if len(resBmi2) != 0 {
			si.CPU.Bmi2 = true
		}

	}
	si.CPU.Threads = uint(runtime.NumCPU())

	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return
	}
	defer f.Close()

	cpu := make(map[string]bool)
	core := make(map[string]bool)

	var cpuID string

	s := bufio.NewScanner(f)
	for s.Scan() {
		if sl := reTwoColumns.Split(s.Text(), 2); sl != nil {
			switch sl[0] {
			case "physical id":
				cpuID = sl[1]
				cpu[cpuID] = true
			case "core id":
				coreID := fmt.Sprintf("%s/%s", cpuID, sl[1])
				core[coreID] = true
			case "vendor_id":
				if si.CPU.Vendor == "" {
					si.CPU.Vendor = sl[1]
				}
			case "model name":
				if si.CPU.Model == "" {
					// CPU model, as reported by /proc/cpuinfo, can be a bit ugly. Clean up...
					model := reExtraSpace.ReplaceAllLiteralString(sl[1], " ")
					si.CPU.Model = strings.Replace(model, "- ", "-", 1)
				}
			case "cache size":
				if si.CPU.Cache == 0 {
					if m := reCacheSize.FindStringSubmatch(sl[1]); m != nil {
						if cache, err := strconv.ParseUint(m[1], 10, 64); err == nil {
							si.CPU.Cache = uint(cache)
						}
					}
				}
			}
		}
	}
	if s.Err() != nil {
		return
	}

	// getNodeInfo() must have run first, to detect if we're dealing with a virtualized CPU! Detecting number of
	// physical processors and/or cores is totally unreliable in virtualized environments, so let's not do it.
	if si.HostName.HostNameStr == "" || si.Hypervisor != "" {
		return
	}

	si.CPU.Cpus = uint(len(cpu))
	si.CPU.Cores = uint(len(core))

}

//GetCPUArch 获取cpu架构
func GetCPUArch() string {
	tmpArch := runtime.GOARCH
	return tmpArch
}