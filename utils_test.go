package main

import (
	"os"
)

func createFile(name string) (*os.File, error) {
	fn, _ := os.Getwd()
	fn += string(os.PathSeparator) + name
	file, err := os.Create(fn)
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return file, nil
}

func checkFileAndRemove(name string) error {
	newFn, err := os.Getwd()
	if err != nil {
		return err
	}
	newFn += string(os.PathSeparator) + name
	_, err = os.Stat(newFn)
	if err != nil {
		return err
	}

	return os.Remove(newFn)
}

func createNestedFoldersWithFiles(dirName, filename string, level int) (string, error) {

	fn, err := os.Getwd()
	if err != nil {
		return "", err
	}

	incrementalFn := fn
	for i := 0; i < level; i++ {
		incrementalFn += string(os.PathSeparator) + dirName
		err = os.Mkdir(incrementalFn, 0755)
		if err != nil {
			return "", err
		}
		f := incrementalFn + string(os.PathSeparator) + filename
		file, err := os.Create(f)
		if err != nil {
			return "", err
		}
		err = file.Close()
		if err != nil {
			return "", err
		}
	}

	return fn + string(os.PathSeparator) + dirName, nil
}

func removeNestedFolder(dirName string) error {
	fn, err := os.Getwd()
	if err != nil {
		return err
	}

	return os.RemoveAll(fn + string(os.PathSeparator) + dirName)
}
