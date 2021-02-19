package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckFolder(t *testing.T) {
	t.Parallel()
	require.Error(t, checkFolder(flags{
		HelpCmd:      nil,
		VersionCmd:   nil,
		Directory:    nil,
		ExtensionCmd: nil,
		ContainsCmd:  nil,
		SnakeCmd:     nil,
		CamelCmd:     nil,
	}))

	here, err := os.Getwd()
	require.NoError(t, err)

	f := flags{
		HelpCmd:      nil,
		VersionCmd:   nil,
		Directory:    &here,
		ExtensionCmd: nil,
		ContainsCmd:  nil,
		SnakeCmd:     nil,
		CamelCmd:     nil,
	}
	require.NoError(t, checkFolder(f))

	require.NoError(t, err)
	folderInvalid := here + "/notExists"
	f.Directory = &folderInvalid
	require.Error(t, checkFolder(f))
}
