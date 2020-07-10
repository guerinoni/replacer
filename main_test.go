package main

import (
	"os/exec"
	"testing"
)

func TestHelpMenu(t *testing.T) {
	cmd := exec.Command("./replacer", "-h")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Error("error")
	}

	sout := string(out)
	if len(sout) == 0 {
		t.Error("output is empty")
	}
}
