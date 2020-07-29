package main

import (
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
var extensionCmd *string

func createFlags() {
	flag.String("v", "", "Return version of replacer.")
	directory = flag.String("d", "", "Specify working directory. (Required)")
	extensionCmd = flag.String("ext", "", "Choose extension to change (<from> <to>).")
}

func exec(extraArgs []string) {
	if err := checkFolder(); err != nil {
		fmt.Println("Folder error")
		os.Exit(1)
	}

	if *extensionCmd != "" {
		execChangeExtension(*directory, *extensionCmd, extraArgs[0])
	}
}

func checkFolder() error {
	fi, err := os.Stat(*directory)
	if fi != nil && err == nil {
		return nil
	}

	return err
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
