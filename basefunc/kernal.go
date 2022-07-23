package basefunc

// Kernel information.
type Kernel struct {
	Release string `json:"release,omitempty"`
	Version string `json:"version,omitempty"`
	// Architecture string `json:"architecture,omitempty"`
}

//getKernelInfo 获取内核信息
func (si *HostInfo) getKernelInfo() {
	si.Kernel.Release = slurpFile("/proc/sys/kernel/osrelease")
	si.Kernel.Version = slurpFile("/proc/sys/kernel/version")

	//*************-syscall包暂时不支持-*****************//
	// var uname syscall.Utsname
	// if err := syscall.Uname(&uname); err != nil {
	// 	return
	// }

	// si.Kernel.Architecture = strings.TrimRight(string((*[65]byte)(unsafe.Pointer(&uname.Machine))[:]), "\000")
}
