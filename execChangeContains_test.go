package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecChangeContains(t *testing.T) {
	t.Parallel()
	file, err := createFile("asdf")
	require.NoError(t, err)
	require.NoError(t, execChangeContains(file.Name(), "sd", "ds"))
	require.NoError(t, checkFileAndRemove("adsf"))
}

func TestExecChangeContainsRecursive(t *testing.T) {
	t.Parallel()
	fn, err := createNestedFoldersWithFiles("TestExecChangeContainsRecursive", "foo.txt", 5)
	require.NoError(t, err)
	defer removeNestedFolder("TestExecChangeContainsRecursive")
	require.NoError(t, execChangeContains(fn, "oo", "xx"))
	require.NoError(t, checkFileInNestedFolder("TestExecChangeContainsRecursive", "fxx.txt"))
}

func TestExecChangeContainsError(t *testing.T) {
	t.Parallel()
	require.Error(t, execChangeContains("", ".txt", ".ttt"))
}

func BenchmarkExecChangeContainsOneFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, _ := createFile("asdf")
		_ = execChangeContains(file.Name(), "sd", "ds")
		_ = checkFileAndRemove("adsf")
	}
}

func BenchmarkExecChangeContainsDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn, _ := createNestedFoldersWithFiles("folder", "foo.txt", 10)
		_ = execChangeContains(fn, "oo", "xx")
		removeNestedFolder("folder")
	}
}

func BenchmarkExecChangeContainsLotDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn, _ := createNestedFoldersWithFiles("folder", "foo.txt", 100)
		_ = execChangeContains(fn, "oo", "xx")
		removeNestedFolder("folder")
	}
}

func BenchmarkExecChangeContainsManyDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn, _ := createNestedFoldersWithFiles("folder", "foo.txt", 1000)
		_ = execChangeContains(fn, "oo", "xx")
		removeNestedFolder("folder")
	}
}
