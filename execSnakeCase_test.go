package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecSnakeCase(t *testing.T) {
	fn, _ := os.Getwd()
	fn += string(os.PathSeparator) + "mainApplication.go"
	file, err := os.Create(fn)
	defer func() {
		err = file.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	if err != nil {
		t.Error(err)
	}

	require.NoError(t, execSnakeCase(file.Name()))

	newFn, _ := os.Getwd()
	newFn += string(os.PathSeparator) + "main_application.go"
	if _, err := os.Stat(newFn); os.IsNotExist(err) {
		t.Errorf("file not exists after change extension")
	}

	os.Remove(newFn)
}

func TestExecSnakeCaseInFolder(t *testing.T) {
	fn, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	fn += string(os.PathSeparator) + "folder"
	err = os.Mkdir(fn, 0777)
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err = os.RemoveAll(fn)
		if err != nil {
			t.Error(err)
		}
	}()

	filename := fn + string(os.PathSeparator) + "camelCase.cpp"
	file, errFile := os.Create(filename)
	if errFile != nil {
		t.Error(errFile)
	}

	err = file.Close()
	if err != nil {
		t.Error(err)
	}

	require.NoError(t, execSnakeCase(file.Name()))

	newFn := fn + string(os.PathSeparator) + "camel_case.cpp"
	if _, err := os.Stat(newFn); os.IsNotExist(err) {
		t.Error(err)
	}
}

func TestExecSnakeCaseAllFolder(t *testing.T) {
	fn, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	fn += string(os.PathSeparator) + "folder"
	err = os.Mkdir(fn, 0777)
	if err != nil {
		t.Error(err)
	}

	defer func() {
		err = os.RemoveAll(fn)
		if err != nil {
			t.Error(err)
		}
	}()

	filename := fn + string(os.PathSeparator) + "camelCase.cpp"
	file, errFile := os.Create(filename)
	if errFile != nil {
		t.Error(errFile)
	}

	err = file.Close()
	if err != nil {
		t.Error(err)
	}

	filename = fn + string(os.PathSeparator) + "camelCase.txt"
	file, errFile = os.Create(filename)
	if errFile != nil {
		t.Error(errFile)
	}

	err = file.Close()
	if err != nil {
		t.Error(err)
	}

	require.NoError(t, execSnakeCase(fn))

	newFn := fn + string(os.PathSeparator) + "camel_case.cpp"
	if _, err := os.Stat(newFn); os.IsNotExist(err) {
		t.Error(err)
	}

	newFn = fn + string(os.PathSeparator) + "camel_case.txt"
	if _, err := os.Stat(newFn); os.IsNotExist(err) {
		t.Error(err)
	}
}

func TestExecSnakeCaseError(t *testing.T) {
	require.Error(t, execSnakeCase(""))
}
