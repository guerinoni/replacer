package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecSnakeCase(t *testing.T) {
	t.Parallel()
	file, err := createFile("mainApplication.go")
	require.NoError(t, err)
	require.NoError(t, execSnakeCase(file.Name()))
	require.NoError(t, checkFileAndRemove("main_application.go"))
}

func TestExecSnakeCaseInFolder(t *testing.T) {
	t.Parallel()
	fn, err := createNestedFoldersWithFiles("folder1", "camelCase.cpp", 1)
	require.NoError(t, err)
	defer removeNestedFolder("folder1")
	require.NoError(t, execSnakeCase(fn))
	require.NoError(t, checkFileInNestedFolder("folder1", "camel_case.cpp"))
}

func TestExecSnakeCaseAllFolder(t *testing.T) {
	t.Parallel()
	fn, err := createNestedFoldersWithFiles("folder2", "camelCase.cpp", 10)
	require.NoError(t, err)
	defer removeNestedFolder("folder2")
	require.NoError(t, execSnakeCase(fn))
	require.NoError(t, checkFileInNestedFolder("folder2", "camel_case.cpp"))
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
	t.Parallel()
	require.Error(t, execSnakeCase(""))
}

func BenchmarkExecSnakeCaseOneFile(b *testing.B) {
	_ = b
	file, _ := createFile("camelCase.cpp")
	_ = execSnakeCase(file.Name())
	_ = checkFileAndRemove("camel_case.cpp")
}

func BenchmarkExecSnakeCaseDir(b *testing.B) {
	_ = b
	fn, _ := createNestedFoldersWithFiles("folder", "camelCase.cpp", 10)
	defer removeNestedFolder("folder")
	_ = execSnakeCase(fn)
}

func BenchmarkExecSnakeCaseLotDir(b *testing.B) {
	_ = b
	fn, _ := createNestedFoldersWithFiles("folder", "camelCase.cpp", 100)
	defer removeNestedFolder("folder")
	_ = execSnakeCase(fn)
}

func BenchmarkExecSnakeCaseManyDir(b *testing.B) {
	_ = b
	fn, _ := createNestedFoldersWithFiles("folder", "camelCase.cpp", 1000)
	defer removeNestedFolder("folder")
	_ = execSnakeCase(fn)
}
