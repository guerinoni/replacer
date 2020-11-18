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

func TestExecChangeExtensionError(t *testing.T) {
	require.Error(t, execChangeExtension("", ".txt", ".ttt"))
}
