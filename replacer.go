package main

import(
	// "fmt"
	"errors"
)

func checkFolder(args []string) error {
	if len(args) <= 1 {
		return errors.New("must contain folder argument")
	}

	return nil
}