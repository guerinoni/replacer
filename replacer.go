package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	cmdExt = "ext"
)

func needHelp(args []string) bool {
	return len(args) > 1 && args[1] == "-h"
}

func printHelp() {
	fmt.Println("	folder (first argument)")
	fmt.Println("	ext <from> 2 <to> (change extension of every file)")
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

func execCmd(args []string) error {
	if len(args) <= 2 {
		return errors.New("command missing")
	}

	switch cmd := args[2]; cmd {
	case cmdExt:
		execChangeExtension(args[1], args[3], args[5])
	default:
		fmt.Printf("command %s not found\n", cmd)
	}

	return nil
}

func execChangeExtension(rootDir string, from string, to string) {
	if !strings.HasPrefix(from, ".") {
		from = "." + from
	}

	if !strings.HasPrefix(to, ".") {
		to = "." + to
	}

	err := filepath.Walk(rootDir, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(info.Name()) == from {
			src := filename
			dst := strings.TrimRight(src, from)
			dst += to
			if err := os.Rename(src, dst); err != nil {
				fmt.Println(err)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("error walking on ", rootDir)
	}
}
