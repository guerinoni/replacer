package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"unicode"
)

func execChangeExtension(rootDir, from, to string) error {
	if !strings.HasPrefix(from, ".") {
		from = "." + from
	}

	if !strings.HasPrefix(to, ".") {
		to = "." + to
	}

	var wg sync.WaitGroup
	err := filepath.Walk(rootDir, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		go func() {
			if filepath.Ext(info.Name()) == from {
				src := filename
				dst := strings.TrimSuffix(src, from)
				dst += to
				if err := os.Rename(src, dst); err != nil {
					fmt.Println(err)
				}
			}
			wg.Done()
		}()

		return nil
	})

	wg.Wait()
	return err
}

func execChangeContains(rootDir, from, to string) error {
	var wg sync.WaitGroup
	err := filepath.Walk(rootDir, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		go func() {
			if strings.Contains(filepath.Base(info.Name()), from) {
				src := filename
				dst := strings.ReplaceAll(src, from, to)
				if err := os.Rename(src, dst); err != nil {
					fmt.Println(err)
				}
			}
			wg.Done()
		}()

		return nil
	})

	wg.Wait()
	return err
}

func execSnakeCase(rootDir string) error {
	var wg sync.WaitGroup
	err := filepath.Walk(rootDir, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			newName := ""
			for _, v := range info.Name() {
				if !unicode.IsUpper(v) {
					newName += string(v)
				} else {
					newName += "_" + string(unicode.ToLower(v))
				}
			}

			newPath := strings.TrimRight(filename, info.Name())
			_ = os.Rename(filename, newPath+newName)
		}()

		return nil
	})

	wg.Wait()
	return err
}
