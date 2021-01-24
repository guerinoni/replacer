package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecCamelCase(t *testing.T) {
	file, err := createFile("main_application.go")
	require.NoError(t, err)
	require.NoError(t, execCamelCase(file.Name()))
	require.NoError(t, checkFileAndRemove("mainApplication.go"))
}

func TestExecCamelCaseInFolder(t *testing.T) {
	fn, err := createNestedFoldersWithFiles("folder", "snake_case.cpp", 1)
	require.NoError(t, err)
	defer removeNestedFolder("folder")
	require.NoError(t, execCamelCase(fn))
	require.NoError(t, checkFileInNestedFolder("folder", "snakeCase.cpp"))
}

func TestExecCamelCaseAllFolder(t *testing.T) {
	fn, err := createNestedFoldersWithFiles("folder", "snake_case.cpp", 10)
	require.NoError(t, err)
	defer removeNestedFolder("folder")
	require.NoError(t, execCamelCase(fn))
	require.NoError(t, checkFileInNestedFolder("folder", "snakeCase.cpp"))
}

func TestExecCamelCaseFileWithSpace(t *testing.T) {
	fn, err := createNestedFoldersWithFiles("folder", "name with spaces.go", 1)
	require.NoError(t, err)
	defer removeNestedFolder("folder")
	require.NoError(t, execCamelCase(fn))
	require.NoError(t, checkFileInNestedFolder("folder", "nameWithSpaces.go"))
}

func TestExecCamelCaseFileWithCapitalLetter(t *testing.T) {
	fn, err := createNestedFoldersWithFiles("folder", "CapitalLetterName.go", 1)
	require.NoError(t, err)
	defer removeNestedFolder("folder")
	require.NoError(t, execCamelCase(fn))
	require.NoError(t, checkFileInNestedFolder("folder", "capitalLetterName.go"))
}

func TestExecCamelCaseError(t *testing.T) {
	require.Error(t, execCamelCase(""))
}

func BenchmarkExecCamelCaseOneFile(b *testing.B) {
	file, _ := createFile("camel_case.cpp")
	_ = execCamelCase(file.Name())
	_ = checkFileAndRemove("camelCase.cpp")
}

func BenchmarkExecCamelCaseDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("folder", "camel_case.cpp", 10)
	defer removeNestedFolder("folder")
	_ = execCamelCase(fn)
}

func BenchmarkExecCamelCaseLotDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("folder", "camel_case.cpp", 100)
	defer removeNestedFolder("folder")
	_ = execCamelCase(fn)
}

func BenchmarkExecCamelCaseManyDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("folder", "camel_case.cpp", 1000)
	defer removeNestedFolder("folder")
	_ = execCamelCase(fn)
}
