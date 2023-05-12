package io

import (
	"os"
	"testing"
)

func TestFileInfo(t *testing.T) {

	stat, err := os.Stat("./fileinfo_test.go")
	if err != nil {
		t.Errorf("Result was incorrect, got error %s.\n", err.Error())
	}

	res := FileInfo(stat)
	if res["cTime"] <= 0 {
		t.Errorf("Result was incorrect, got %d, expected > 0.\n", res["cTime"])
	}

	if res["mTime"] <= 0 {
		t.Errorf("Result was incorrect, got %d, expected > 0.\n", res["mTime"])
	}

	if res["aTime"] <= 0 {
		t.Errorf("Result was incorrect, got %d, expected > 0.\n", res["aTime"])
	}
}
