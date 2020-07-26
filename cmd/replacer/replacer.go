package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	cmdExt      = "ext"
	cmdContains = "contains"
)

var directory *string

func createFlags() {
	directory = flag.String("d", "", "Specify working directory. (Required)")
	flag.String("ext", "", "Choose extension to change (<from> <to>).")
}

func exec() {
	if err := checkFolder(); err != nil {
		fmt.Println("Folder error")
		os.Exit(1)
	}
}

func checkFolder() error {
	fi, err := os.Stat(*directory)
	if fi != nil && err == nil {
		return nil
	}

	return err
}

func execCmd(args []string) error {
	if len(args) <= 2 {
		return errors.New("command missing")
	}

	switch cmd := args[2]; cmd {
	case cmdExt:
		execChangeExtension(args[1], args[3], args[5])
	case cmdContains:
		execChangeContains(args[1], args[3], args[5])
	default:
		fmt.Printf("command %s not found\n", cmd)
	}

	return nil
}

func execChangeExtension(rootDir, from, to string) {
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

func execChangeContains(rootDir, oldStr, newStr string) {
	err := filepath.Walk(rootDir, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.Contains(filepath.Base(info.Name()), oldStr) {
			src := filename
			dst := strings.ReplaceAll(src, oldStr, newStr)
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
