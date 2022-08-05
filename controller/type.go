package controller

type ResData struct {
	Code  string `json:"code"`
	Error string `json:"error,omitempty"`
	Data  interface{}
}

//EnvConf yaml配置的key不能用 _ 否则会信息丢失
type EnvConf struct {
	Role       string     `json:"role"`
	MasterIP   string     `json:"masterip"`
	MasterPort string     `json:"masterport"`
	Mode       string     `json:"mode"`
	Host       []HostPara `json:"host"`
}

type HostPara struct {
	IP string `json:"ip"`
	// User     string `json:"user"`    //暂时只支持root
	Password string `json:"password"`
	Port     string `json:"port,omitempty"`
}
