package main

import (
	"os"
	"testing"
)

func TestExecChangeExtensionWithDot(t *testing.T) {
	file, err := os.Create("foo.txt")
	if err != nil {
		t.Errorf("error creating file")
	}

	execChangeExtension(file.Name(), ".txt", ".c")

	if _, err := os.Stat("foo.c"); os.IsNotExist(err) {
		t.Errorf("file not exists after change extension")
	}

	os.Remove("foo.c")
}

func TestExecChangeExtensionWithoutDot(t *testing.T) {
	file, err := os.Create("foo.txt")
	if err != nil {
		t.Errorf("error creating file")
	}

	execChangeExtension(file.Name(), "txt", "c")

	if _, err := os.Stat("foo.c"); os.IsNotExist(err) {
		t.Errorf("file not exists after change extension")
	}

	os.Remove("foo.c")
}
