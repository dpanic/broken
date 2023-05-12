// +build linux

package io

import (
	"os"
	"syscall"
)

// FileInfo gets modification and creation time of a os.FileInfo
func FileInfo(val os.FileInfo) (res map[string]int64) {
	res = make(map[string]int64)

	stat := val.Sys().(*syscall.Stat_t)

	cTime := stat.Ctim.Sec
	res["cTime"] = cTime

	mTime := stat.Mtim.Sec
	res["mTime"] = mTime

	aTime := stat.Atim.Sec
	res["aTime"] = aTime

	return
}
