package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecChangeContains(t *testing.T) {
	file, err := createFile("asdf")
	require.NoError(t, err)
	require.NoError(t, execChangeContains(file.Name(), "sd", "ds"))
	require.NoError(t, checkFileAndRemove("adsf"))
}

func TestExecChangeContainsRecursive(t *testing.T) {
	fn, err := createNestedFoldersWithFiles("folder", "foo.txt", 5)
	require.NoError(t, err)
	defer removeNestedFolder("folder")
	require.NoError(t, execChangeContains(fn, "oo", "xx"))
	require.NoError(t, checkFileInNestedFolder("folder", "fxx.txt"))
}

func TestExecChangeContainsError(t *testing.T) {
	require.Error(t, execChangeContains("", ".txt", ".ttt"))
}

func BenchmarkExecChangeContainsOneFile(b *testing.B) {
	file, _ := createFile("asdf")
	_ = execChangeContains(file.Name(), "sd", "ds")
	_ = checkFileAndRemove("adsf")
}

func BenchmarkExecChangeContainsDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("folder", "foo.txt", 10)
	defer removeNestedFolder("folder")
	_ = execChangeContains(fn, "oo", "xx")
}

func BenchmarkExecChangeContainsLotDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("folder", "foo.txt", 100)
	defer removeNestedFolder("folder")
	_ = execChangeContains(fn, "oo", "xx")
}

func BenchmarkExecChangeContainsManyDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("folder", "foo.txt", 1000)
	defer removeNestedFolder("folder")
	_ = execChangeContains(fn, "oo", "xx")
}
