package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func isDir(file *os.File) bool {
	_, err := file.ReadDir(0)
	return err == nil
}

func makeAsm(codeWriter *codeWriter, file *os.File) {
	parser := newParser(file)
	defer parser.close()

	codeWriter.setFileName(file.Name())

	for {
		if !parser.hasMoreCommands() {
			break
		}

		parser.advance()
		cm := parser.command

		if len(cm) > 0 {
			switch parser.commandType() {
			case C_ARITHMETIC:
				codeWriter.writeArthmethic(parser.commandItself())
			case C_PUSH, C_POP:
				index, err := strconv.Atoi(parser.arg2())
				if err != nil {
					panic(err)
				}
				codeWriter.writePushPop(parser.commandItself(), parser.arg1(), index)
			case C_LABEL:
				codeWriter.writeLabel(parser.arg1())
			case C_GOTO:
				codeWriter.writeGoto(parser.arg1())
			case C_IF:
				codeWriter.writeIf(parser.arg1())
			case C_FUNCTION:
				numLocal, err := strconv.Atoi(parser.arg2())
				if err != nil {
					panic(err)
				}
				codeWriter.writeFunction(parser.arg1(), numLocal)
			case C_RETURN:
				codeWriter.writeReturn()
			case C_CALL:
				numLocal, err := strconv.Atoi(parser.arg2())
				if err != nil {
					panic(err)
				}
				codeWriter.writeCall(parser.arg1(), numLocal)
			}
		}
	}
}

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	codeWriter := newCodeWriter(file)
	defer codeWriter.close()

	codeWriter.writeBootload()

	if isDir(file) {
		file.Seek(0, io.SeekStart)
		files, err := file.ReadDir(0)
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			path := filepath.Join(filePath, file.Name())

			// 拡張子チェック
			if filepath.Ext(file.Name()) != ".vm" {
				log.Printf("%s is not vm file", file.Name())
				continue
			}
			f, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			makeAsm(codeWriter, f)
		}
		return
	}
	makeAsm(codeWriter, file)
}
