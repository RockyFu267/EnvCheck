package basefunc

//ChartYaml chart下的Chart.yaml
type ChartYaml struct {
	Name        string                 `json:"name"`
	Home        string                 `json:"home"`
	Version     string                 `json:"version"`
	AppVersion  string                 `json:"appVersion"`
	APIVersion  string                 `json:"apiVersion"`
	Description string                 `json:"description"`
	Sources     []string               `json:"sources"`
	Maintainers []ChartYamlMaintainers `json:"maintainers"`
	Email       string                 `json:"  email"`
}

//ChartYamlMaintainers ChartYaml-Maintainers
type ChartYamlMaintainers struct {
	Name string `json:"name"`
}

type TestTest struct {
	Aaa string      `json:"aaa,omitempty"`
	Bbb string      `json:"bbb,omitempty"`
	Ccc TestTestCcc `json:"ccc,omitempty"`
	Ddd string      `json:"ddd,omitempty"`
}
type TestTestCcc struct {
	Cc1c TestTestCccCc2c `json:"cc1c,omitempty"`
}

type TestTestCccCc2c struct {
	Cc1c2c string `json:"cc1c2c,omitempty"`
}
