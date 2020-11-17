package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func execChangeExtension(rootDir, from, to string) error {
	if !strings.HasPrefix(from, ".") {
		from = "." + from
	}

	if !strings.HasPrefix(to, ".") {
		to = "." + to
	}

	err := filepath.Walk(rootDir, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
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

	return err
}

func execChangeContains(rootDir, from, to string) error {
	err := filepath.Walk(rootDir, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
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

	return err
}

func execSnakeCase(rootDir string) error {
	err := filepath.Walk(rootDir, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fInfo, err := os.Stat(rootDir)
		if err != nil {
			return err
		}

		basePath := rootDir
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

	return err
}
