package controller

type ResData struct {
	Code  string `json:"code"`
	Error string `json:"error,omitempty"`
	Data  interface{}
}

//EnvConf yaml配置的key不能用 _ 否则会信息丢失
type EnvConf struct {
	MasterIP     string     `json:"masterip"`
	MasterPort   string     `json:"masterport"`
	Mode         string     `json:"mode"`
	RemotePath   string     `json:"remotepath"`
	DiskTestInfo DiskTest   `json:"disktest"`
	Host         []HostPara `json:"host"`
}
type DiskTest struct {
	Type string `json:"type"`
	//DiskTestPath 不能为空
	DiskTestPath string `json:"disktestpath"`
}

type HostPara struct {
	IP           string `json:"ip"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Port         string `json:"sshport,omitempty"`
	DiskTestBool bool   `json:"disktest,omitempty"`
	DiskTestPath string `json:"disktestpath,omitempty"`
}
