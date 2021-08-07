package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecSnakeCase(t *testing.T) {
	file, err := createFile("TestExecSnakeCase.go")
	require.NoError(t, err)
	require.NoError(t, execSnakeCase(file.Name()))
	require.NoError(t, checkFileAndRemove("test_exec_snake_case.go"))
}

func TestExecSnakeCaseFromLower(t *testing.T) {
	file, err := createFile("test exec snake case from lower.go")
	require.NoError(t, err)
	require.NoError(t, execSnakeCase(file.Name()))
	require.NoError(t, checkFileAndRemove("test_exec_snake_case_from_lower.go"))
}

func TestExecSnakeCaseInFolder(t *testing.T) {
	fn, err := createNestedFoldersWithFiles("TestExecSnakeCaseInFolder", "camelCase.cpp", 1)
	require.NoError(t, err)

	defer removeNestedFolder("TestExecSnakeCaseInFolder")
	require.NoError(t, execSnakeCase(fn))
	require.NoError(t, checkFileInNestedFolder("TestExecSnakeCaseInFolder", "camel_case.cpp"))
}

func TestExecSnakeCaseAllFolder(t *testing.T) {
	fn, err := createNestedFoldersWithFiles("TestExecSnakeCaseAllFolder", "camelCase.cpp", 10)
	require.NoError(t, err)

	defer removeNestedFolder("TestExecSnakeCaseAllFolder")
	require.NoError(t, execSnakeCase(fn))
	require.NoError(t, checkFileInNestedFolder("TestExecSnakeCaseAllFolder", "camel_case.cpp"))
}

func TestExecSnakeCaseFileWithSpace(t *testing.T) {
	fn, err := createNestedFoldersWithFiles("TestExecSnakeCaseFileWithSpace", "name with spaces.go", 1)
	require.NoError(t, err)

	defer removeNestedFolder("TestExecSnakeCaseFileWithSpace")
	require.NoError(t, execSnakeCase(fn))
	require.NoError(t, checkFileInNestedFolder("TestExecSnakeCaseFileWithSpace", "name_with_spaces.go"))
}

func TestExecSnakeCaseFileWithCapitalLetter(t *testing.T) {
	fn, err := createNestedFoldersWithFiles("TestExecSnakeCaseFileWithCapitalLetter", "CapitalLetterName.go", 1)
	require.NoError(t, err)

	defer removeNestedFolder("TestExecSnakeCaseFileWithCapitalLetter")
	require.NoError(t, execSnakeCase(fn))
	require.NoError(t, checkFileInNestedFolder("TestExecSnakeCaseFileWithCapitalLetter", "capital_letter_name.go"))
}

func TestExecSnakeCaseError(t *testing.T) {
	require.Error(t, execSnakeCase(""))
}

func BenchmarkExecSnakeCaseOneFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, _ := createFile("camelCase.cpp")
		_ = execSnakeCase(file.Name())
		_ = checkFileAndRemove("camel_case.cpp")
	}
}

func BenchmarkExecSnakeCaseDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn, _ := createNestedFoldersWithFiles("folder", "camelCase.cpp", 10)

		defer removeNestedFolder("folder")

		_ = execSnakeCase(fn)
	}
}

func BenchmarkExecSnakeCaseLotDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn, _ := createNestedFoldersWithFiles("folder", "camelCase.cpp", 100)
		_ = execSnakeCase(fn)

		removeNestedFolder("folder")
	}
}

func BenchmarkExecSnakeCaseManyDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn, _ := createNestedFoldersWithFiles("folder", "camelCase.cpp", 1000)
		_ = execSnakeCase(fn)

		removeNestedFolder("folder")
	}
}
