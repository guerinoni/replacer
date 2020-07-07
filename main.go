package main

import(
	"fmt"
	"os"
)

func main() {
	fmt.Println("replacer version 0.0.1")
	if err := checkFolder(os.Args); err != nil {
		fmt.Println("Error parsing folder ->", err)
	}
}