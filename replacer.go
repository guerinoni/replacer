package main

import (
	"errors"
	"fmt"
	"os"
)

func needHelp(args []string) bool {
	return len(args) > 1 && args[1] == "-h"
}

func printHelp() {
	fmt.Println("	folder (first argument)")
	fmt.Println("	ext <from> -> <to> (change extension of every file)")
}

func checkFolder(args []string) error {
	if len(args) <= 1 {
		return errors.New("must contain folder argument")
	}

	if _, err := os.Stat(args[1]); os.IsNotExist(err) {
		return errors.New("this folder not exists")
	}

	return nil
}
