package main

import (
	"os"
	"path/filepath"
)

func isDir(file *os.File) bool {
	_, err := file.ReadDir(0)
	if err == nil {
		return true
	}
	return false
}

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	if isDir(file) {
		files, err := file.ReadDir(0)
		if err != nil {
			panic(err)
		}

		codeWriter := newCodeWriter(file)
		defer codeWriter.close()

		for _, file := range files {
			d := filepath.Join(filePath, file.Name())
			if file.Type().IsDir() {
				panic("can't translate directory")
			}
			f, err := os.Open(d)
			if err != nil {
				panic(err)
			}

			parser := newParser(f)
			defer parser.close()
		}
		return
	}

	codeWriter := newCodeWriter(file)
	defer codeWriter.close()
	parser := newParser(file)
	defer parser.close()
}
