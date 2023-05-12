// +build darwin

package io

import (
	"os"
	"syscall"
	"time"
)

// FileInfo gets modification and creation time of a os.FileInfo
func FileInfo(val os.FileInfo) (res map[string]int64) {
	res = make(map[string]int64)

	stat := val.Sys().(*syscall.Stat_t)

	uts := stat.Ctimespec
	cTime := time.Unix(int64(uts.Sec), int64(uts.Nsec))
	res["cTime"] = cTime.Unix()

	uts = stat.Mtimespec
	mTime := time.Unix(int64(uts.Sec), int64(uts.Nsec))
	res["mTime"] = mTime.Unix()

	uts = stat.Atimespec
	aTime := time.Unix(int64(uts.Sec), int64(uts.Nsec))
	res["aTime"] = aTime.Unix()

	return
}
