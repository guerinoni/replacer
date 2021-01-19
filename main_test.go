package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckFolder(t *testing.T) {
	t.Parallel()
	require.Error(t, checkFolder(flags{}))

	folder := "/home/"
	f := flags{
		Directory:    &folder,
		ExtensionCmd: nil,
		ContainsCmd:  nil,
		SnakeCmd:     nil,
	}
	require.NoError(t, checkFolder(f))

	folderInvalid := "/home/asdf/"
	f.Directory = &folderInvalid
	require.Error(t, checkFolder(f))
}
