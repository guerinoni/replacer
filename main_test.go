package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCaseExec struct {
	name  string
	f     flags
	extra []string
}

func makeTestCaseExec(name string, extraArgs []string, f flags) testCaseExec {
	return testCaseExec{
		name:  name,
		extra: extraArgs,
		f:     f,
	}
}

//nolint:funlen
func TestExec(t *testing.T) {
	t.Parallel()
	tmpDir := func() *string {
		tmp, err := os.MkdirTemp(os.TempDir(), "fake")
		require.NoError(t, err)

		return &tmp
	}

	newBool := func() *bool {
		b := true

		return &b
	}

	newString := func() *string {
		b := "lol"

		return &b
	}

	testCases := []testCaseExec{
		makeTestCaseExec("noting", []string{}, flags{
			HelpCmd:      new(bool),
			VersionCmd:   new(bool),
			Directory:    tmpDir(),
			ExtensionCmd: new(string),
			ContainsCmd:  new(string),
			SnakeCmd:     new(bool),
			CamelCmd:     new(bool),
		}),

		makeTestCaseExec("suggestion of dir missing", []string{}, flags{
			HelpCmd:      new(bool),
			VersionCmd:   new(bool),
			Directory:    new(string),
			ExtensionCmd: new(string),
			ContainsCmd:  new(string),
			SnakeCmd:     new(bool),
			CamelCmd:     new(bool),
		}),

		makeTestCaseExec("version", []string{}, flags{
			HelpCmd:      new(bool),
			VersionCmd:   newBool(),
			Directory:    new(string),
			ExtensionCmd: new(string),
			ContainsCmd:  new(string),
			SnakeCmd:     new(bool),
			CamelCmd:     new(bool),
		}),

		makeTestCaseExec("help", []string{}, flags{
			HelpCmd:      newBool(),
			VersionCmd:   new(bool),
			Directory:    new(string),
			ExtensionCmd: new(string),
			ContainsCmd:  new(string),
			SnakeCmd:     new(bool),
			CamelCmd:     new(bool),
		}),

		makeTestCaseExec("snake", []string{}, flags{
			HelpCmd:      new(bool),
			VersionCmd:   new(bool),
			Directory:    tmpDir(),
			ExtensionCmd: new(string),
			ContainsCmd:  new(string),
			SnakeCmd:     newBool(),
			CamelCmd:     new(bool),
		}),

		makeTestCaseExec("camel", []string{}, flags{
			HelpCmd:      new(bool),
			VersionCmd:   new(bool),
			Directory:    tmpDir(),
			ExtensionCmd: new(string),
			ContainsCmd:  new(string),
			SnakeCmd:     new(bool),
			CamelCmd:     newBool(),
		}),

		makeTestCaseExec("extension", []string{"txt"}, flags{
			HelpCmd:      new(bool),
			VersionCmd:   new(bool),
			Directory:    tmpDir(),
			ExtensionCmd: newString(),
			ContainsCmd:  new(string),
			SnakeCmd:     new(bool),
			CamelCmd:     new(bool),
		}),

		makeTestCaseExec("extension", []string{"txt"}, flags{
			HelpCmd:      new(bool),
			VersionCmd:   new(bool),
			Directory:    tmpDir(),
			ExtensionCmd: new(string),
			ContainsCmd:  newString(),
			SnakeCmd:     new(bool),
			CamelCmd:     new(bool),
		}),
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := exec(&tt.f, tt.extra)
			require.NoError(t, err)
		})
	}
}
