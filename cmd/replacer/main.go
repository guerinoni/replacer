package main

import (
	"flag"
	"os"
)

var version string

func main() {
	createFlags()

	if len(os.Args) == 1 {
		flag.PrintDefaults()
	}

	flag.Parse()

	exec()
}
