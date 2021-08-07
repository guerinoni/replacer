package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func createFile(name string) (*os.File, error) {
	fn, _ := os.Getwd()
	fn += string(os.PathSeparator) + name
	file, err := os.Create(fn)

	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}

	err = file.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close file: %w", err)
	}

	return file, nil
}

func checkFileAndRemove(name string) error {
	newFn, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current dir: %w", err)
	}

	newFn += string(os.PathSeparator) + name
	_, err = os.Stat(newFn)

	if err != nil {
		return fmt.Errorf("failed check info of file: %w", err)
	}

	err = os.Remove(newFn)
	if err != nil {
		return fmt.Errorf("failed remove file: %w", err)
	}

	return nil
}

func createNestedFoldersWithFiles(dirName, filename string, level int) (string, error) {
	fn, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to insert user: %w", err)
	}

	incrementalFn := fn
	for i := 0; i < level; i++ {
		incrementalFn += string(os.PathSeparator) + dirName
		err = os.Mkdir(incrementalFn, 0o755)

		if err != nil {
			return "", fmt.Errorf("failed to make dir: %w", err)
		}

		f := incrementalFn + string(os.PathSeparator) + filename
		file, err := os.Create(f)

		if err != nil {
			return "", fmt.Errorf("failed to create file: %w", err)
		}

		err = file.Close()

		if err != nil {
			return "", fmt.Errorf("failed to close file: %w", err)
		}
	}

	return fn + string(os.PathSeparator) + dirName, nil
}

func removeNestedFolder(dirName string) {
	fn, err := os.Getwd()
	if err != nil {
		return
	}

	os.RemoveAll(fn + string(os.PathSeparator) + dirName)
}

func checkFileInNestedFolder(dir, file string) error {
	err := filepath.Walk(dir, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if info.Name() != file {
			return fmt.Errorf("file not found: %s", file)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk path: %w", err)
	}

	return nil
}
