package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecChangeExtensionWithDot(t *testing.T) {
	file, err := createFile("foo.txt")
	require.NoError(t, err)
	require.NoError(t, execChangeExtension(file.Name(), ".txt", ".ttt"))
	require.NoError(t, checkFileAndRemove("foo.ttt"))
}

func TestExecChangeExtensionWithoutDot(t *testing.T) {
	file, err := createFile("foo.txt")
	require.NoError(t, err)
	require.NoError(t, execChangeExtension(file.Name(), "txt", "ttt"))
	require.NoError(t, checkFileAndRemove("foo.ttt"))
}

func TestExecChangeExtensionWithNameEqualExtension(t *testing.T) {
	file, err := createFile("foo.foo")
	require.NoError(t, err)
	require.NoError(t, execChangeExtension(file.Name(), ".foo", ".ttt"))
	require.NoError(t, checkFileAndRemove("foo.ttt"))
}

func TestExecChangeExtensionRecursive(t *testing.T) {
	fn, err := createNestedFoldersWithFiles("folder", "foo.txt", 2)
	require.NoError(t, err)
	require.NoError(t, execChangeExtension(fn, ".txt", ".ttt"))
	require.NoError(t, removeNestedFolder("folder"))
}

func TestExecChangeExtensionError(t *testing.T) {
	require.Error(t, execChangeExtension("", ".txt", ".ttt"))
}

func BenchmarkExecChangeExtensionOneFile(b *testing.B) {
	file, _ := createFile("foo.txt")
	execChangeExtension(file.Name(), ".txt", ".ttt")
	checkFileAndRemove("foo.ttt")
}

func BenchmarkExecChangeExtensionDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("dir", "file.txt", 10)
	execChangeExtension(fn, ".txt", ".ttt")
	removeNestedFolder("dir")
}

func BenchmarkExecChangeExtensionLotDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("dir", "file.txt", 100)
	execChangeExtension(fn, ".txt", ".ttt")
	removeNestedFolder("dir")
}

func BenchmarkExecChangeExtensionManyDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("dir", "file.txt", 1000)
	execChangeExtension(fn, ".txt", ".ttt")
	removeNestedFolder("dir")
}
