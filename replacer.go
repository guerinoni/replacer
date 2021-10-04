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

	err := filepath.WalkDir(rootDir, func(filename string, info os.DirEntry, err error) error {
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
				if err := os.Rename(filename, filename[0:len(filename)-len(from)]+to); err != nil {
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

	err := filepath.WalkDir(rootDir, func(filename string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			if strings.Contains(info.Name(), from) {
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

	err := filepath.WalkDir(rootDir, func(filename string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			uppers := 0
			relativeName := info.Name()
			for i, v := range relativeName {
				if i > 0 && unicode.IsUpper(v) {
					uppers++
				}
			}
			res := make([]byte, len(filename)+uppers)
			copy(res, filename[0:len(filename)-len(relativeName)])
			ptr := len(filename) - len(relativeName)
			for i, v := range relativeName {
				if unicode.IsSpace(v) {
					res[ptr] = '_'
					ptr++
					continue
				}

				if unicode.IsUpper(v) {
					if i > 0 {
						res[ptr] = '_'
						ptr++
					}
					res[ptr] = byte(unicode.ToLower(v))
					ptr++
				} else {
					res[ptr] = byte(v)
					ptr++
				}
			}
			_ = os.Rename(filename, string(res))
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

	err := filepath.WalkDir(rootDir, func(fileName string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			forceUpperNext := false

			removed := 0
			relativeName := info.Name()
			for _, v := range relativeName {
				if v == '_' || v == '-' || unicode.IsSpace(v) {
					removed++
				}
			}

			res := make([]byte, len(fileName)-removed)
			ptr := len(fileName) - len(relativeName)
			copy(res, fileName[0:ptr])

			for i, v := range relativeName {
				if forceUpperNext {
					res[ptr] = byte(unicode.ToUpper(v))
					ptr++
					forceUpperNext = false
					continue
				}

				if i == 0 {
					res[ptr] = byte(unicode.ToLower(v))
					ptr++
					continue
				}

				if v == '_' || v == '-' || unicode.IsSpace(v) {
					forceUpperNext = true
				} else {
					res[ptr] = byte(v)
					ptr++
				}
			}
			_ = os.Rename(fileName, string(res))
		}()

		return nil
	})

	wg.Wait()

	if err != nil {
		return fmt.Errorf("failed to walk path: %w", err)
	}

	return nil
}
