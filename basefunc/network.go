package basefunc

import (
	bc "EnvCheck/basecmd"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// NetworkDevice information.
type NetworkDevice struct {
	Name       string `json:"name,omitempty"`
	MACAddress string `json:"macaddress,omitempty"`
	Speed      string `json:"speed,omitempty"` // device max supported speed in Mbps
}

//getNetworkInfo 获取网卡信息
func (si *HostInfo) getNetworkInfo() {
	sysClassNet := "/sys/class/net"
	devices, err := ioutil.ReadDir(sysClassNet)
	if err != nil {
		return
	}

	si.Network = make([]NetworkDevice, 0)
	for _, link := range devices {
		fullpath := path.Join(sysClassNet, link.Name())
		dev, err := os.Readlink(fullpath)
		if err != nil {
			continue
		}

		if strings.HasPrefix(dev, "../../devices/virtual/") {
			continue
		}

		device := NetworkDevice{
			Name:       link.Name(),
			MACAddress: slurpFile(path.Join(fullpath, "address")),
			Speed:      getMaxSpeed(link.Name()),
		}

		si.Network = append(si.Network, device)
	}
}

//getMaxSpeed 获取网卡速度
func getMaxSpeed(input string) (res string) {
	cmdTmp := `ethtool ` + input + ` | grep -i speed | awk '{print $2}'`
	//是否存在英伟达显卡
	resTMP, err := bc.CmdAndChangeDirToResAllInOne("./", cmdTmp)
	if err != nil {
		log.Println("Get NetCard error: ", err)
		res = "unknown"
		return
	}
	if len(resTMP) == 0 {
		res = "unknown"
		return
	}

	return resTMP[0]
}
