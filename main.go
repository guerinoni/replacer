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
