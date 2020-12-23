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
	fn, err := createNestedFoldersWithFiles("folder", "foo.txt", 10)
	require.NoError(t, err)
	defer removeNestedFolder("folder")
	require.NoError(t, execChangeExtension(fn, ".txt", ".ttt"))
	require.NoError(t, checkFileInNestedFolder("folder", "foo.ttt"))
}

func TestExecChangeExtensionError(t *testing.T) {
	require.Error(t, execChangeExtension("", ".txt", ".ttt"))
}

func BenchmarkExecChangeExtensionOneFile(b *testing.B) {
	file, _ := createFile("foo.txt")
	_ = execChangeExtension(file.Name(), ".txt", ".ttt")
	_ = checkFileAndRemove("foo.ttt")
}

func BenchmarkExecChangeExtensionDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("dir", "file.txt", 10)
	_ = execChangeExtension(fn, ".txt", ".ttt")
	removeNestedFolder("dir")
}

func BenchmarkExecChangeExtensionLotDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("dir", "file.txt", 100)
	_ = execChangeExtension(fn, ".txt", ".ttt")
	removeNestedFolder("dir")
}

func BenchmarkExecChangeExtensionManyDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("dir", "file.txt", 1000)
	_ = execChangeExtension(fn, ".txt", ".ttt")
	removeNestedFolder("dir")
}
