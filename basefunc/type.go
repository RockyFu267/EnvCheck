package basefunc

//OSInfo 主机信息
type HostInfo struct {
	OperatingSystem string `json:"operatingsystem"`
	CPUARCH         string `json:"cpuarch"`
	HostName        string `json:"hostname"`
}
