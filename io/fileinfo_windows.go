// +build windows

package io

import (
	"os"
	"syscall"
	"time"
)

// FileInfo gets modification and creation time of a os.FileInfo
func FileInfo(val os.FileInfo) (res map[string]int64) {
	res = make(map[string]int64)

	stat := val.Sys().(*syscall.Win32FileAttributeData)

	uts := time.Unix(0, stat.CreationTime.Nanoseconds())
	res["cTime"] = uts.Unix()

	uts = time.Unix(0, stat.LastAccessTime.Nanoseconds())
	res["aTime"] = uts.Unix()

	uts = time.Unix(0, stat.LastWriteTime.Nanoseconds())
	res["mTime"] = uts.Unix()

	return
}
