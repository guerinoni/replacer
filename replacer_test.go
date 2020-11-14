package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecChangeExtensionWithDot(t *testing.T) {
	fn, _ := os.Getwd()
	fn += string(os.PathSeparator) + "foo.txt"
	file, err := os.Create(fn)
	defer func() {
		err = file.Close()
		require.NoError(t, err)
	}()

	require.NoError(t, err)
	require.NoError(t, execChangeExtension(file.Name(), ".txt", ".ttt"))

	newFn, _ := os.Getwd()
	newFn += string(os.PathSeparator) + "foo.ttt"
	_, err = os.Stat(newFn)
	require.False(t, os.IsNotExist(err))

	os.Remove(newFn)
}

func TestExecChangeExtensionWithoutDot(t *testing.T) {
	fn, _ := os.Getwd()
	fn += string(os.PathSeparator) + "foo.txt"
	file, err := os.Create(fn)
	defer func() {
		err = file.Close()
		require.NoError(t, err)
	}()

	require.NoError(t, err)
	require.NoError(t, execChangeExtension(file.Name(), "txt", "ttt"))

	newFn, _ := os.Getwd()
	newFn += string(os.PathSeparator) + "foo.ttt"
	_, err = os.Stat(newFn)
	require.False(t, os.IsNotExist(err))
	os.Remove(newFn)
}

func TestExecChangeExtensionWithNameEqualExtension(t *testing.T) {
	fn, _ := os.Getwd()
	fn += string(os.PathSeparator) + "foo.foo"
	file, err := os.Create(fn)
	defer func() {
		err = file.Close()
		require.NoError(t, err)
	}()
	require.NoError(t, err)
	require.NoError(t, execChangeExtension(file.Name(), ".foo", ".ttt"))

	newFn, _ := os.Getwd()
	newFn += string(os.PathSeparator) + "foo.ttt"
	_, err = os.Stat(newFn)
	require.False(t, os.IsNotExist(err))
	os.Remove(newFn)
}

func TestExecChangeExtensionRecursive(t *testing.T) {
	fn, err := os.Getwd()
	require.NoError(t, err)

	fn += string(os.PathSeparator) + "folder"
	err = os.Mkdir(fn, 0777)
	require.NoError(t, err)

	firstFolderCreate, _ := os.Getwd()
	firstFolderCreate += string(os.PathSeparator) + "folder"
	defer func() {
		err = os.RemoveAll(firstFolderCreate)
		require.NoError(t, err)
	}()

	filename := fn + string(os.PathSeparator) + "foo.txt"
	file, err := os.Create(filename)
	require.NoError(t, err)

	err = file.Close()
	require.NoError(t, err)

	fn += string(os.PathSeparator) + "folder1"
	err = os.Mkdir(fn, 0777)
	require.NoError(t, err)

	filename1 := fn + string(os.PathSeparator) + "foo.txt"
	file1, err := os.Create(filename1)
	require.NoError(t, err)

	err = file1.Close()
	require.NoError(t, err)

	require.NoError(t, execChangeExtension(firstFolderCreate, ".txt", ".ttt"))

	here, _ := os.Getwd()
	newFn := here + string(os.PathSeparator) + "folder" + string(os.PathSeparator) + "foo.ttt"
	_, err = os.Stat(newFn)
	require.False(t, os.IsNotExist(err))

	newFn1 := here + string(os.PathSeparator) + "folder" + string(os.PathSeparator) + "foo.ttt"
	_, err = os.Stat(newFn1)
	require.False(t, os.IsNotExist(err))
}

func TestChangeExtensionError(t *testing.T) {
	require.Error(t, execChangeExtension("", ".txt", ".ttt"))
}
func TestExecChangeContains(t *testing.T) {
	fn, _ := os.Getwd()
	fn += string(os.PathSeparator) + "asdf"
	file, err := os.Create(fn)
	defer func() {
		err = file.Close()
		if err != nil {
			t.Error(err)
		}
	}()

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
	defer func() {
		err = file.Close()
		if err != nil {
			t.Error(err)
		}
	}()

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

	execSnakeCase(file.Name())

	newFn := fn + string(os.PathSeparator) + "camel_case.cpp"
	if _, err := os.Stat(newFn); os.IsNotExist(err) {
		t.Error(err)
	}
}

func TestExecSnakeAllFolder(t *testing.T) {
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

	execSnakeCase(fn)

	newFn := fn + string(os.PathSeparator) + "camel_case.cpp"
	if _, err := os.Stat(newFn); os.IsNotExist(err) {
		t.Error(err)
	}

	newFn = fn + string(os.PathSeparator) + "camel_case.txt"
	if _, err := os.Stat(newFn); os.IsNotExist(err) {
		t.Error(err)
	}
}
