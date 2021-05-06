package main

import (
	"flag"
	"fmt"
)

var version string

func main() {
	flags := newFlags()

	flag.Parse()

	err := exec(flags, flag.Args())
	if err != nil {
		fmt.Println(err)
	}
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

func exec(f *flags, extraArgs []string) error {
	if *f.VersionCmd {
		fmt.Println("replacer version: ", version)

		return nil
	}

	if *f.HelpCmd {
		flag.PrintDefaults()

		return nil
	}

	if *f.Directory == "" {
		fmt.Println("missing -d <folder>")

		return nil
	}

	if *f.SnakeCmd {
		err := execSnakeCase(*f.Directory)

		return err
	}

	if *f.CamelCmd {
		err := execCamelCase(*f.Directory)

		return err
	}

	if *f.ExtensionCmd != "" {
		err := execChangeExtension(*f.Directory, *f.ExtensionCmd, extraArgs[0])

		return err
	}

	if *f.ContainsCmd != "" {
		err := execChangeContains(*f.Directory, *f.ContainsCmd, extraArgs[0])

		return err
	}

	return nil
}
