package basefunc

//OSInfo 主机信息
type HostInfo struct {
	// CPUARCH         string          `json:"cpuarch,omitempty"`     //暂时不用这个方法
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
	CheckRes    bool   `json:"CheckRes"`
}
type Storage struct {
	RootDirSize string          `json:"rootdirsize,omitempty"`
	DiskInfo    []StorageDevice `json:"storageDevice"`
}

type Others struct {
	TimeZone string `json:"timezone,omitempty"`
}
