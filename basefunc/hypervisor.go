package basefunc

import (
	"io/ioutil"
	"strings"
	"unsafe"

	"EnvCheck/cpuid"
)

// https://en.wikipedia.org/wiki/CPUID#EAX.3D0:_Get_vendor_ID
var hvmap = map[string]string{
	"bhyve bhyve ": "bhyve",
	"KVMKVMKVM":    "kvm",
	"Microsoft Hv": "hyperv",
	" lrpepyh vr":  "parallels",
	"VMwareVMware": "vmware",
	"XenVMMXenVMM": "xenhvm",
}

//isHypervisorActive
func isHypervisorActive() bool {
	var info [4]uint32
	cpuid.CPUID(&info, 0x1)
	return info[2]&(1<<31) != 0
}

//getHypervisorCpuid
func getHypervisorCpuid(ax uint32) string {
	var info [4]uint32
	cpuid.CPUID(&info, ax)
	return hvmap[strings.TrimRight(string((*[12]byte)(unsafe.Pointer(&info[1]))[:]), "\000")]
}

//getHypervisor 获取底层信息
func (si *HostInfo) getHypervisor() {
	if !isHypervisorActive() {
		if hypervisorType := slurpFile("/sys/hypervisor/type"); hypervisorType != "" {
			if hypervisorType == "xen" {
				si.Hypervisor = "xenpv"
			}
		}
		return
	}

	// KVM has been caught to move its real signature to this leaf, and put something completely different in the
	// standard location. So this leaf must be checked first.
	if hv := getHypervisorCpuid(0x40000100); hv != "" {
		si.Hypervisor = hv
		return
	}

	if hv := getHypervisorCpuid(0x40000000); hv != "" {
		si.Hypervisor = hv
		return
	}

	// getBIOSInfo() must have run first, to detect BIOS vendor
	if slurpFile("/sys/class/dmi/id/bios_vendor") == "Bochs" {
		si.Hypervisor = "bochs"
		return
	}

	si.Hypervisor = "unknown"
}

//slurpFile Read one-liner text files, strip newline.
func slurpFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}
