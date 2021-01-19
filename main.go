package main

import (
	"flag"
	"fmt"
	"os"
)

var version string

func main() {
	flags := createFlags()

	if len(os.Args) > 1 && os.Args[1] == "-v" {
		fmt.Println("replacer version: ", version)
		os.Exit(0)
	}

	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()

	exec(flags, flag.Args())
}

type flags struct {
	Directory    *string
	ExtensionCmd *string
	ContainsCmd  *string
	SnakeCmd     *bool
	CamelCmd     *bool
}

func createFlags() (f flags) {
	flag.String("v", "", "Return version of replacer.")
	f.Directory = flag.String("d", "", "Specify working directory. (Required)")
	f.ExtensionCmd = flag.String("ext", "", "Choose extension to change <from> <to>. (i.e. replacer -d . -ext txt cpp")
	f.ContainsCmd = flag.String("contains", "",
		"Choose substr to change <from> <to>. (i.e. replacer -d . -contains as ss)")
	f.SnakeCmd = flag.Bool("snake", false, "Rename all files in path specified with snake case. (i.e. replacer -d . -snake)")
	f.CamelCmd = flag.Bool("camel", false, "Raname all files in specified path with camel case. (i.e replacer -d . -camel)")

	return
}

func exec(f flags, extraArgs []string) {
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
	fi, err := os.Stat(*f.Directory)
	if fi != nil && err == nil {
		return nil
	}

	return err
}
