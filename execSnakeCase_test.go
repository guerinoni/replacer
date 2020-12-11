package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecSnakeCase(t *testing.T) {
	file, err := createFile("mainApplication.go")
	require.NoError(t, err)
	require.NoError(t, execSnakeCase(file.Name()))
	require.NoError(t, checkFileAndRemove("main_application.go"))
}

func TestExecSnakeCaseInFolder(t *testing.T) {
	fn, err := createNestedFoldersWithFiles("folder", "camelCase.cpp", 1)
	require.NoError(t, err)
	defer removeNestedFolder("folder")
	require.NoError(t, execSnakeCase(fn))
	require.NoError(t, checkFileInNestedFolder("folder", "camel_case.cpp"))
}

func TestExecSnakeCaseAllFolder(t *testing.T) {
	fn, err := createNestedFoldersWithFiles("folder", "camelCase.cpp", 10)
	require.NoError(t, err)
	defer removeNestedFolder("folder")
	require.NoError(t, execSnakeCase(fn))
	require.NoError(t, checkFileInNestedFolder("folder", "camel_case.cpp"))
}

func TestExecSnakeCaseError(t *testing.T) {
	require.Error(t, execSnakeCase(""))
}

func BenchmarkExecSnakeCaseOneFile(b *testing.B) {
	file, _ := createFile("camelCase.cpp")
	execSnakeCase(file.Name())
	checkFileAndRemove("camel_case.cpp")
}

func BenchmarkExecSnakeCaseDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("folder", "camelCase.cpp", 10)
	defer removeNestedFolder("folder")
	execSnakeCase(fn)
}

func BenchmarkExecSnakeCaseLotDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("folder", "camelCase.cpp", 100)
	defer removeNestedFolder("folder")
	execSnakeCase(fn)
}

func BenchmarkExecSnakeCaseManyDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("folder", "camelCase.cpp", 1000)
	defer removeNestedFolder("folder")
	execSnakeCase(fn)
}
