package basefunc

// Version of the EnvCheck library.
const Version = "0.2.1"

var HostInfoList []HostInfo
var CheckWarning []string
var PostInfoList []PostInfo
var LenHostList int

//OSInfo 主机信息
type HostInfo struct {
	Meta       Meta            `json:"meta"`
	CPU        CPU             `json:"cpu"`
	HostName   HostName        `json:"hostname"`
	Hypervisor string          `json:"hypervisor,omitempty"`
	OS         OS              `json:"os"`
	Kernel     Kernel          `json:"kernel"`
	Memory     Memory          `json:"memorysize"`
	GPU        GPU             `json:"gpu"`
	Network    []NetworkDevice `json:"network"`
	Storage    Storage         `json:"storage"`
	Others     Others          `json:"others"`
}

type HostName struct {
	HostNameStr string `json:"hostname,omitempty"`
	CheckRes    bool   `json:"check_res"`
}
type Storage struct {
	RootDirSize string          `json:"root_dir_size,omitempty"`
	DfInfo      []string        `json:"dfinfo,omitempty"`
	LsblkInfo   []string        `json:"lsblkinfo,omitempty"`
	RaidInfo    bool            `json:"raidinfo,omitempty"`
	DataDir     DataDir         `json:"data_dir"`
	DiskInfo    []StorageDevice `json:"storagedevice"`
	DiskIO      FioResInfo      `json:"diskio"`
}

type DataDir struct {
	MountExistCheck bool   `json:"mount_exist_check"`
	DataDirSize     string `json:"data_dir_size,omitempty"`
	Type            string `json:"type,omitempty"`
}
type Others struct {
	TimeZone     string       `json:"timezone,omitempty"`
	Firewall     Firewall     `json:"firewall"`
	ServiceCheck ServiceCheck `json:"service_check"`
}

type Firewall struct {
	Stop    bool `json:"stop"`
	Disable bool `json:"disable"`
}

type PostInfo struct {
	PostIP string `json:"post_ip"`
	Res    bool   `json:"res"`
}
