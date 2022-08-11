package basefunc

import (
	bc "EnvCheck/basecmd"
	_ "bufio"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

// StorageDevice information.
type StorageDevice struct {
	Name   string `json:"name,omitempty"`
	Size   uint   `json:"size,omitempty"` // device size in GB
	Type   string `json:"type,omitempty"`
	Driver string `json:"driver,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	Model  string `json:"model,omitempty"`
}

//getStorageInfo  获取主机磁盘信息
func (si *HostInfo) getStorageInfo() {
	//获取< / >大小
	si.getRootDirSize()
	//获取是否存在 /data
	si.getDataDir()

	sysBlock := "/sys/block"
	devices, err := ioutil.ReadDir(sysBlock)
	if err != nil {
		return
	}

	si.Storage.DiskInfo = make([]StorageDevice, 0)
	for _, link := range devices {
		fullpath := path.Join(sysBlock, link.Name())
		//加入获取Type的逻辑，判断是SSD还是HHD
		var typeTmp string = "unknown"
		cmdType := `lsblk -d -o name,rota | grep ` + link.Name() + ` | awk '{print $2}'`
		resType, err := bc.CmdAndChangeDirToResAllInOne("./", cmdType)
		if err != nil {
			log.Println("Get resType error: ", err)
		} else {
			if len(resType) != 0 {
				if resType[0] == "1" {
					typeTmp = "HDD"
				}
				if resType[0] == "0" {
					typeTmp = "SSD"
				}
			}
		}
		dev, err := os.Readlink(fullpath)
		if err != nil {
			continue
		}

		if strings.HasPrefix(dev, "../devices/virtual/") {
			continue
		}

		// We could filter all removable devices here, but some systems boot from USB flash disks, and then we
		// would filter them, too. So, let's filter only floppies and CD/DVD devices, and see how it pans out.
		if strings.HasPrefix(dev, "../devices/platform/floppy") || slurpFile(path.Join(fullpath, "device", "type")) == "5" {
			continue
		}

		device := StorageDevice{
			Name:  link.Name(),
			Model: slurpFile(path.Join(fullpath, "device", "model")),
			Type:  typeTmp,
		}

		if driver, err := os.Readlink(path.Join(fullpath, "device", "driver")); err == nil {
			device.Driver = path.Base(driver)
		}

		if vendor := slurpFile(path.Join(fullpath, "device", "vendor")); !strings.HasPrefix(vendor, "0x") {
			device.Vendor = vendor
		}

		size, _ := strconv.ParseUint(slurpFile(path.Join(fullpath, "size")), 10, 64)
		device.Size = uint(size) / 1953125 // GiB

		si.Storage.DiskInfo = append(si.Storage.DiskInfo, device)
	}
}

//getRootDirSize 获取 / 大小
func (si *HostInfo) getRootDirSize() {
	resLsblk, err := bc.CmdAndChangeDirToResAllInOne("./", "lsblk | grep /")
	if err != nil {
		log.Println("Get RootDirSize error: ", err)
		return
	}
	//找出 / 对应的结果
	var resTmp string
	for _, v := range resLsblk {
		if v[len(v)-1:] == "/" {
			resTmp = v
			break
		}
	}
	resTmp = DeleteExtraSpace(resTmp)
	res := strings.Fields(resTmp)
	si.Storage.RootDirSize = res[3]
}

//getDataDir 获取 /data 是否存在挂载点，并获取响应的信息
func (si *HostInfo) getDataDir() {
	//检查是否存在挂载点
	resLsblk, err := bc.CmdAndChangeDirToResAllInOne("./", "lsblk | grep /data")
	if err != nil {
		log.Println("Get DataDirInfo error: ", err)
		return
	}
	if len(resLsblk) == 0 {
		si.Storage.DataDir.MountExistCheck = false
	}
	//找出是否存在 /data
	var resTmp string = ""
	for _, v := range resLsblk {
		if v[len(v)-5:] == "/data" {
			resTmp = v
			break
		}
	}
	if resTmp == "" {
		si.Storage.DataDir.MountExistCheck = false
		return
	} else {
		si.Storage.DataDir.MountExistCheck = true
	}
	//获取的大小
	resTmp = DeleteExtraSpace(resTmp)
	res := strings.Fields(resTmp)
	si.Storage.DataDir.DataDirSize = res[3]

	//检查挂载的类型
	resDfTh, err := bc.CmdAndChangeDirToResAllInOne("./", "df -Th | grep /data")
	if err != nil {
		log.Println("Get DataDirtType error: ", err)
		return
	}
	if len(resDfTh) == 0 {
		si.Storage.DataDir.Type = "unknown"
	}
	//找出是否存在 /data
	var resType string = ""
	for _, v := range resDfTh {
		if v[len(v)-5:] == "/data" {
			resType = v
			break
		}
	}
	if resType == "" {
		si.Storage.DataDir.Type = "unknown"
		return
	}
	resType = DeleteExtraSpace(resType)
	resTypeList := strings.Fields(resType)
	si.Storage.DataDir.Type = resTypeList[1]

}
