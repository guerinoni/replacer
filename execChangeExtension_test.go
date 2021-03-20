package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecChangeExtensionWithDot(t *testing.T) {
	t.Parallel()
	file, err := createFile("foo.txt")
	require.NoError(t, err)
	require.NoError(t, execChangeExtension(file.Name(), ".txt", ".ttt"))
	require.NoError(t, checkFileAndRemove("foo.ttt"))
}

func TestExecChangeExtensionWithoutDot(t *testing.T) {
	t.Parallel()
	file, err := createFile("doo.txt")
	require.NoError(t, err)
	require.NoError(t, execChangeExtension(file.Name(), "txt", "ttt"))
	require.NoError(t, checkFileAndRemove("doo.ttt"))
}

func TestExecChangeExtensionWithNameEqualExtension(t *testing.T) {
	t.Parallel()
	file, err := createFile("eoo.eoo")
	require.NoError(t, err)
	require.NoError(t, execChangeExtension(file.Name(), ".eoo", ".ttt"))
	require.NoError(t, checkFileAndRemove("eoo.ttt"))
}

func TestExecChangeExtensionRecursive(t *testing.T) {
	t.Parallel()
	fn, err := createNestedFoldersWithFiles("TestExecChangeExtensionRecursive", "foo.txt", 10)
	require.NoError(t, err)
	defer removeNestedFolder("TestExecChangeExtensionRecursive")
	require.NoError(t, execChangeExtension(fn, ".txt", ".ttt"))
	require.NoError(t, checkFileInNestedFolder("TestExecChangeExtensionRecursive", "foo.ttt"))
}

func TestExecChangeExtensionError(t *testing.T) {
	t.Parallel()
	require.Error(t, execChangeExtension("", ".txt", ".ttt"))
}

func BenchmarkExecChangeExtensionOneFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, _ := createFile("foo.txt")
		_ = execChangeExtension(file.Name(), ".txt", ".ttt")
		_ = checkFileAndRemove("foo.ttt")
	}
}

func BenchmarkExecChangeExtensionDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn, _ := createNestedFoldersWithFiles("dir", "file.txt", 10)
		_ = execChangeExtension(fn, ".txt", ".ttt")
		removeNestedFolder("dir")
	}
}

func BenchmarkExecChangeExtensionLotDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn, _ := createNestedFoldersWithFiles("dir", "file.txt", 100)
		_ = execChangeExtension(fn, ".txt", ".ttt")
		removeNestedFolder("dir")
	}
}

func BenchmarkExecChangeExtensionManyDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn, _ := createNestedFoldersWithFiles("dir", "file.txt", 1000)
		_ = execChangeExtension(fn, ".txt", ".ttt")
		removeNestedFolder("dir")
	}
}
