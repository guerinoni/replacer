package main

import (
	"os"
	"testing"
)

func TestExecChangeExtensionWithDot(t *testing.T) {
	fn, _ := os.Getwd()
	fn += string(os.PathSeparator) + "foo.txt"
	file, err := os.Create(fn)
	if err != nil {
		t.Errorf("error creating file")
	}

	file.Close()

	execChangeExtension(file.Name(), ".txt", ".ttt")

	newFn, _ := os.Getwd()
	newFn += string(os.PathSeparator) + "foo.ttt"
	if _, err := os.Stat(newFn); os.IsNotExist(err) {
		t.Errorf("file not exists after change extension")
	}

	os.Remove(newFn)
}

func TestExecChangeExtensionWithoutDot(t *testing.T) {
	fn, _ := os.Getwd()
	fn += string(os.PathSeparator) + "foo.txt"
	file, err := os.Create(fn)
	if err != nil {
		t.Errorf("error creating file")
	}

	file.Close()

	execChangeExtension(file.Name(), "txt", "ttt")

	newFn, _ := os.Getwd()
	newFn += string(os.PathSeparator) + "foo.ttt"
	if _, err := os.Stat(newFn); os.IsNotExist(err) {
		t.Errorf("file not exists after change extension")
	}

	os.Remove(newFn)
}
