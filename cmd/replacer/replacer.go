package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

var directory *string
var extensionCmd *string
var containsCmd *string
var snakeCmd *string

func createFlags() {
	flag.String("v", "", "Return version of replacer.")
	directory = flag.String("d", "", "Specify working directory. (Required)")
	extensionCmd = flag.String("ext", "", "Choose extension to change (<from> <to>).")
	containsCmd = flag.String("contains", "", "Choose substr to change (<from> <to>).")
	snakeCmd = flag.String("snake", "", "Rename all files in path specified with snake case.")
}

func exec(extraArgs []string) {
	if *snakeCmd != "" {
		execSnakeCase(*snakeCmd)
		return
	}

	if err := checkFolder(); err != nil {
		fmt.Println("Folder error")
		os.Exit(1)
	}

	if *extensionCmd != "" {
		execChangeExtension(*directory, *extensionCmd, extraArgs[0])
	}

	if *containsCmd != "" {
		execChangeContains(*directory, *containsCmd, extraArgs[0])
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
			dst := strings.TrimSuffix(src, from)
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

func execChangeContains(rootDir, from, to string) {
	err := filepath.Walk(rootDir, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.Contains(filepath.Base(info.Name()), from) {
			src := filename
			dst := strings.ReplaceAll(src, from, to)
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

func execSnakeCase(rootDir string) {
	err := filepath.Walk(rootDir, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		fInfo, err := os.Stat(rootDir)
		if err != nil {
			return err
		}

		var basePath string
		if !fInfo.IsDir() {
			basePath = filepath.Dir(rootDir)
		}

		newName := ""
		for _, v := range info.Name() {
			if !unicode.IsUpper(v) {
				newName += string(v)
			} else {
				newName += "_" + string(unicode.ToLower(v))
			}
		}

		err = os.Rename(filename, basePath+string(os.PathSeparator)+newName)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("error walking on ", rootDir)
	}
}
