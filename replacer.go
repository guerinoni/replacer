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
			defer wg.Done()
			if filepath.Ext(info.Name()) == from {
				src := filename
				dst := strings.TrimSuffix(src, from)
				dst += to
				if err := os.Rename(src, dst); err != nil {
					fmt.Fprintf(os.Stderr, "error: %s", err)
				}
			}
		}()

		return nil
	})

	wg.Wait()

	if err != nil {
		return fmt.Errorf("failed to walk path: %w", err)
	}

	return nil
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
			defer wg.Done()
			if strings.Contains(filepath.Base(info.Name()), from) {
				src := filename
				dst := strings.ReplaceAll(src, from, to)
				if err := os.Rename(src, dst); err != nil {
					fmt.Fprintf(os.Stderr, "error: %s", err)
				}
			}
		}()

		return nil
	})

	wg.Wait()

	if err != nil {
		return fmt.Errorf("failed to walk path: %w", err)
	}

	return nil
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
			for i, v := range info.Name() {
				if unicode.IsSpace(v) {
					newName += "_"

					continue
				}

				if unicode.IsUpper(v) {
					if i > 0 {
						newName += "_"
					}
					newName += string(unicode.ToLower(v))
				} else {
					newName += string(v)
				}
			}

			newPath := strings.TrimRight(filename, info.Name())
			_ = os.Rename(filename, newPath+newName)
		}()

		return nil
	})

	wg.Wait()

	if err != nil {
		return fmt.Errorf("failed to walk path: %w", err)
	}

	return nil
}

func execCamelCase(rootDir string) error {
	var wg sync.WaitGroup

	err := filepath.Walk(rootDir, func(fileName string, info os.FileInfo, err error) error {
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
			forceUpperNext := false
			for i, v := range info.Name() {
				if forceUpperNext {
					newName += string(unicode.ToUpper(v))
					forceUpperNext = false

					continue
				}

				if i == 0 {
					newName += string(unicode.ToLower(v))

					continue
				}

				if string(v) == "_" || string(v) == "-" || unicode.IsSpace(v) {
					forceUpperNext = true
				} else {
					newName += string(v)
				}
			}

			filePath := strings.TrimRight(fileName, info.Name())
			_ = os.Rename(fileName, filePath+newName)
		}()

		return nil
	})

	wg.Wait()

	if err != nil {
		return fmt.Errorf("failed to walk path: %w", err)
	}

	return nil
}
