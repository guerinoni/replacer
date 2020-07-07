package main

import (
	"fmt"
	"os"
)

var version string

func main() {
	if needHelp(os.Args) {
		fmt.Printf("replacer (version %s)\n", version)
		printHelp()
		return
	}

	if err := checkFolder(os.Args); err != nil {
		fmt.Println("Error parsing folder ->", err)
	}
}
