package basefunc

//OSInfo 主机信息
type HostInfo struct {
	OperatingSystem string   `json:"operatingsystem,omitempty"`
	CPUARCH         string   `json:"cpuarch,omitempty"`
	HostName        HostName `json:"hostname"`
	Hypervisor      string   `json:"hypervisor,omitempty"`
	OS              OS       `json:"os"`
}

type HostName struct {
	HostNameStr string `json:"cpuarch,omitempty"`
	CheckRes    bool   `json:"CheckRes"`
}
