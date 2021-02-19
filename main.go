package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var version string

func main() {
	flags := newFlags()

	flag.Parse()

	exec(flags, flag.Args())
}

type flags struct {
	HelpCmd      *bool
	VersionCmd   *bool
	Directory    *string
	ExtensionCmd *string
	ContainsCmd  *string
	SnakeCmd     *bool
	CamelCmd     *bool
}

const (
	directoryCmdDescr = "Specify working directory. (Required)"
	extensionCmdDescr = "Choose extension to change <from> <to>. (i.e. replacer -d . -ext txt cpp"
	containsCmdDescr  = "Choose substr to change <from> <to>. (i.e. replacer -d . -contains as ss)"
	snakeCmdDescr     = "Rename all files in path specified with snake case. (i.e. replacer -d . -snake)"
	camelCmdDescr     = "Raname all files in specified path with camel case. (i.e replacer -d . -camel)"
)

func newFlags() *flags {
	return &flags{
		HelpCmd:      flag.Bool("h", false, "Print all options available."),
		VersionCmd:   flag.Bool("v", false, "Return version of replacer."),
		Directory:    flag.String("d", "", directoryCmdDescr),
		ExtensionCmd: flag.String("ext", "", extensionCmdDescr),
		ContainsCmd:  flag.String("contains", "", containsCmdDescr),
		SnakeCmd:     flag.Bool("snake", false, snakeCmdDescr),
		CamelCmd:     flag.Bool("camel", false, camelCmdDescr),
	}
}

func exec(f *flags, extraArgs []string) {
	if *f.VersionCmd {
		fmt.Println("replacer version: ", version)

		return
	}

	if *f.HelpCmd {
		flag.PrintDefaults()

		return
	}

	if *f.SnakeCmd {
		err := execSnakeCase(*f.Directory)
		if err != nil {
			panic(err)
		}

		return
	}

	if *f.CamelCmd {
		err := execCamelCase(*f.Directory)
		if err != nil {
			panic(err)
		}

		return
	}

	if *f.ExtensionCmd != "" {
		err := execChangeExtension(*f.Directory, *f.ExtensionCmd, extraArgs[0])
		if err != nil {
			panic(err)
		}
	}

	if *f.ContainsCmd != "" {
		err := execChangeContains(*f.Directory, *f.ContainsCmd, extraArgs[0])
		if err != nil {
			panic(err)
		}
	}
}

func checkFolder(f flags) error {
	if f.Directory == nil {
		return errors.New("directory var is not valid")
	}

	fi, err := os.Stat(*f.Directory)
	if fi != nil && err == nil {
		return nil
	}

	return err
}
