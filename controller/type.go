package controller

type ResData struct {
	Code  string `json:"code"`
	Error string `json:"error,omitempty"`
	Data  interface{}
}

type EnvConf struct {
	Role       string     `json:"role"`
	MasterIP   string     `json:"master_ip"`
	MasterPort string     `json:"master_port"`
	Mode       string     `json:"mode"`
	Host       []HostPara `json:"host"`
}

type HostPara struct {
	IP       string `json:"ip"`
	User     string `json:"user"`
	Password string `json:"password"`
}
