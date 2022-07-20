// Package cpuid gives Go programs access to CPUID opcode.
package cpuid

// CPUID returns processor identification and feature information.
func CPUID(info *[4]uint32, ax uint32)
