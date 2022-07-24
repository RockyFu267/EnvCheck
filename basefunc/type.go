package basefunc

//OSInfo 主机信息
type HostInfo struct {
	OperatingSystem string   `json:"operatingsystem,omitempty"`
	CPUARCH         string   `json:"cpuarch,omitempty"`
	CPU             CPU      `json:"cpu"`
	HostName        HostName `json:"hostname"`
	Hypervisor      string   `json:"hypervisor,omitempty"`
	OS              OS       `json:"os"`
	Kernel          Kernel   `json:"kernel"`
	Memory          Memory   `json:"memorysize"`
}

type HostName struct {
	HostNameStr string `json:"cpuarch,omitempty"`
	CheckRes    bool   `json:"CheckRes"`
}
