package basefunc

type OSInfo struct {
	OperatingSystem string `json:"operatingsystem"`
	CPUARCH         string `json:"cpuarch"`
	HostName        string `json:"hostname"`
}
