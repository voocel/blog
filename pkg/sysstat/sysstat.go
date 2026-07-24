// Package sysstat exposes lightweight host metrics (disk, memory) for the admin dashboard.
package sysstat

import (
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

// clampPercent rounds to the nearest integer and constrains the result to 0-100.
func clampPercent(v float64) int {
	p := int(v + 0.5)
	if p < 0 {
		return 0
	}
	if p > 100 {
		return 100
	}
	return p
}

// DiskUsagePercent returns the used disk percentage (0-100) of the filesystem
// containing path. Returns 0 when the value cannot be determined.
func DiskUsagePercent(path string) int {
	if path == "" {
		path = "."
	}
	u, err := disk.Usage(path)
	if err != nil || u == nil {
		return 0
	}
	return clampPercent(u.UsedPercent)
}

// MemoryUsagePercent returns the used system memory percentage (0-100).
// Returns 0 when the value cannot be determined.
func MemoryUsagePercent() int {
	vm, err := mem.VirtualMemory()
	if err != nil || vm == nil {
		return 0
	}
	return clampPercent(vm.UsedPercent)
}
