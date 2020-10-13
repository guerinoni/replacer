package main

import (
	"flag"
	"fmt"
	"os"
)

var version string

func main() {
	createFlags()

	if len(os.Args) > 1 && os.Args[1] == "-v" {
		fmt.Println("replacer version: ", version)
		os.Exit(0)
	}

	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()

	exec(flag.Args())
}

var directory *string
var extensionCmd *string
var containsCmd *string
var snakeCmd *string

func createFlags() {
	flag.String("v", "", "Return version of replacer.")
	directory = flag.String("d", "", "Specify working directory. (Required)")
	extensionCmd = flag.String("ext", "", "Choose extension to change (<from> <to>).")
	containsCmd = flag.String("contains", "", "Choose substr to change (<from> <to>).")
	snakeCmd = flag.String("snake", "", "Rename all files in path specified with snake case.")
}

func exec(extraArgs []string) {
	if *snakeCmd != "" {
		execSnakeCase(*snakeCmd)
		return
	}

	if err := checkFolder(); err != nil {
		fmt.Println("Folder error")
		os.Exit(1)
	}

	if *extensionCmd != "" {
		execChangeExtension(*directory, *extensionCmd, extraArgs[0])
	}

	if *containsCmd != "" {
		execChangeContains(*directory, *containsCmd, extraArgs[0])
	}
}

func checkFolder() error {
	fi, err := os.Stat(*directory)
	if fi != nil && err == nil {
		return nil
	}

	return err
}
