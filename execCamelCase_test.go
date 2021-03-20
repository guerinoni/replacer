package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecCamelCase(t *testing.T) {
	t.Parallel()
	file, err := createFile("test_exec_camel_case.go")
	require.NoError(t, err)
	require.NoError(t, execCamelCase(file.Name()))
	require.NoError(t, checkFileAndRemove("testExecCamelCase.go"))
}

func TestExecCamelCaseInFolder(t *testing.T) {
	t.Parallel()
	fn, err := createNestedFoldersWithFiles("TestExecCamelCaseInFolder", "snake_case.cpp", 1)
	require.NoError(t, err)
	defer removeNestedFolder("TestExecCamelCaseInFolder")
	require.NoError(t, execCamelCase(fn))
	require.NoError(t, checkFileInNestedFolder("TestExecCamelCaseInFolder", "snakeCase.cpp"))
}

func TestExecCamelCaseAllFolder(t *testing.T) {
	t.Parallel()
	fn, err := createNestedFoldersWithFiles("TestExecCamelCaseAllFolder", "snake_case.cpp", 10)
	require.NoError(t, err)
	defer removeNestedFolder("TestExecCamelCaseAllFolder")
	require.NoError(t, execCamelCase(fn))
	require.NoError(t, checkFileInNestedFolder("TestExecCamelCaseAllFolder", "snakeCase.cpp"))
}

func TestExecCamelCaseFileWithSpace(t *testing.T) {
	t.Parallel()
	fn, err := createNestedFoldersWithFiles("TestExecCamelCaseFileWithSpace", "name with spaces.go", 1)
	require.NoError(t, err)
	defer removeNestedFolder("TestExecCamelCaseFileWithSpace")
	require.NoError(t, execCamelCase(fn))
	require.NoError(t, checkFileInNestedFolder("TestExecCamelCaseFileWithSpace", "nameWithSpaces.go"))
}

func TestExecCamelCaseFileWithCapitalLetter(t *testing.T) {
	t.Parallel()
	fn, err := createNestedFoldersWithFiles("TestExecCamelCaseFileWithCapitalLetter", "CapitalLetterName.go", 1)
	require.NoError(t, err)
	defer removeNestedFolder("TestExecCamelCaseFileWithCapitalLetter")
	require.NoError(t, execCamelCase(fn))
	require.NoError(t, checkFileInNestedFolder("TestExecCamelCaseFileWithCapitalLetter", "capitalLetterName.go"))
}

func TestExecCamelCaseError(t *testing.T) {
	t.Parallel()
	require.Error(t, execCamelCase(""))
}

func BenchmarkExecCamelCaseOneFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, _ := createFile("camel_case.cpp")
		_ = execCamelCase(file.Name())
		_ = checkFileAndRemove("camelCase.cpp")
	}
}

func BenchmarkExecCamelCaseDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn, _ := createNestedFoldersWithFiles("folder", "camel_case.cpp", 10)
		_ = execCamelCase(fn)
		removeNestedFolder("folder")
	}
}

func BenchmarkExecCamelCaseLotDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn, _ := createNestedFoldersWithFiles("folder", "camel_case.cpp", 100)
		_ = execCamelCase(fn)
		removeNestedFolder("folder")
	}
}

func BenchmarkExecCamelCaseManyDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn, _ := createNestedFoldersWithFiles("folder", "camel_case.cpp", 1000)
		_ = execCamelCase(fn)
		removeNestedFolder("folder")
	}
}
