package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	ignore := []string{".git"}
	files := []string{}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	files = appendFilesInDir(files, ignore, wd)

	for _, file := range files {
		fmt.Println(file)
	}
}

func appendFilesInDir(files, ignore []string, dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		name := entry.Name()

		if slices.Contains(ignore, name) {
			continue
		}

		path := dir + "/" + name

		if entry.IsDir() {
			files = appendFilesInDir(files, ignore, path)
		} else {
			files = append(files, path)
		}
	}

	return files
}
