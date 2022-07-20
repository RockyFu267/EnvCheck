package basefunc

//OSInfo 主机信息
type HostInfo struct {
	OperatingSystem string `json:"operatingsystem,omitempty"`
	CPUARCH         string `json:"cpuarch,omitempty"`
	HostName        string `json:"hostname,omitempty"`
	Hypervisor      string `json:"hypervisor,omitempty"`
	OS              OS     `json:"os"`
}
