package main

import(
	"os"
	"errors"
)

func checkFolder(args []string) error {
	if len(args) <= 1 {
		return errors.New("must contain folder argument")
	}

	if _, err := os.Stat(args[1]); os.IsNotExist(err) {
		return errors.New("this folder not exists")
	}

	return nil
}