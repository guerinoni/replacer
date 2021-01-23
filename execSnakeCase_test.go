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

func TestExecSnakeCaseFileWithSpace(t *testing.T) {
	fn, err := createNestedFoldersWithFiles("folder", "name with spaces.go", 1)
	require.NoError(t, err)
	defer removeNestedFolder("folder")
	require.NoError(t, execSnakeCase(fn))
	require.NoError(t, checkFileInNestedFolder("folder", "name_with_spaces.go"))
}

func TestExecSnakeCaseFileWithCapitalLetter(t *testing.T) {
	fn, err := createNestedFoldersWithFiles("folder", "CapitalLetterName.go", 1)
	require.NoError(t, err)
	defer removeNestedFolder("folder")
	require.NoError(t, execSnakeCase(fn))
	require.NoError(t, checkFileInNestedFolder("folder", "capital_letter_name.go"))
}

func TestExecSnakeCaseError(t *testing.T) {
	require.Error(t, execSnakeCase(""))
}

func BenchmarkExecSnakeCaseOneFile(b *testing.B) {
	file, _ := createFile("camelCase.cpp")
	_ = execSnakeCase(file.Name())
	_ = checkFileAndRemove("camel_case.cpp")
}

func BenchmarkExecSnakeCaseDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("folder", "camelCase.cpp", 10)
	defer removeNestedFolder("folder")
	_ = execSnakeCase(fn)
}

func BenchmarkExecSnakeCaseLotDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("folder", "camelCase.cpp", 100)
	defer removeNestedFolder("folder")
	_ = execSnakeCase(fn)
}

func BenchmarkExecSnakeCaseManyDir(b *testing.B) {
	fn, _ := createNestedFoldersWithFiles("folder", "camelCase.cpp", 1000)
	defer removeNestedFolder("folder")
	_ = execSnakeCase(fn)
}
