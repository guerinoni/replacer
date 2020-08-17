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

func TestExecChangeExtensionWithNameEqualExtension(t *testing.T) {
	fn, _ := os.Getwd()
	fn += string(os.PathSeparator) + "foo.foo"
	file, err := os.Create(fn)
	defer file.Close()
	if err != nil {
		t.Errorf("error creating file")
	}

	execChangeExtension(file.Name(), ".foo", ".ttt")

	newFn, _ := os.Getwd()
	newFn += string(os.PathSeparator) + "foo.ttt"
	if _, err := os.Stat(newFn); os.IsNotExist(err) {
		t.Errorf("file not exists after change extension")
	}

	os.Remove(newFn)
}

func TestExecChangeContains(t *testing.T) {
	fn, _ := os.Getwd()
	fn += string(os.PathSeparator) + "asdf"
	file, err := os.Create(fn)
	defer file.Close()
	if err != nil {
		t.Errorf("error creating file")
	}

	execChangeContains(file.Name(), "sd", "ds")

	newFn, _ := os.Getwd()
	newFn += string(os.PathSeparator) + "adsf"
	if _, err := os.Stat(newFn); os.IsNotExist(err) {
		t.Errorf("file not exists after change extension")
	}

	os.Remove(newFn)
}

func TestExecSnakeCase(t *testing.T) {
	fn, _ := os.Getwd()
	fn += string(os.PathSeparator) + "mainApplication.go"
	file, err := os.Create(fn)
	defer file.Close()
	if err != nil {
		t.Error(err)
	}

	execSnakeCase(file.Name())

	newFn, _ := os.Getwd()
	newFn += string(os.PathSeparator) + "main_application.go"
	if _, err := os.Stat(newFn); os.IsNotExist(err) {
		t.Errorf("file not exists after change extension")
	}

	os.Remove(newFn)
}
