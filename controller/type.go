package controller

type ResData struct {
	Code  string `json:"code"`
	Error string `json:"error,omitempty"`
	Data  interface{}
}

//EnvConf yaml配置的key不能用 _ 否则会信息丢失
type EnvConf struct {
	MasterIP   string     `json:"masterip"`
	MasterPort string     `json:"masterport"`
	Mode       string     `json:"mode"`
	RemotePath string     `json:"remotepath"`
	Disk       DiskPara   `json:"disk"`
	Host       []HostPara `json:"host"`
}
type DiskPara struct {
	Type string `json:"type"`
	//DiskTestPath 不能为空
	DiskCheckPath string `json:"diskcheckpath"`
}

type HostPara struct {
	IP            string `json:"ip"`
	User          string `json:"user"`
	Password      string `json:"password"`
	Port          string `json:"sshport,omitempty"`
	DiskCheck     string `json:"diskcheck,omitempty"`
	DiskCheckPath string `json:"diskcheckpath,omitempty"`
}
