package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func isDir(file *os.File) bool {
	_, err := file.ReadDir(0)
	return err == nil
}

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if isDir(file) {
		file.Seek(0, 0)
		files, err := file.ReadDir(0)
		if err != nil {
			panic(err)
		}

		codeWriter := newCodeWriter(file)
		defer codeWriter.close()

		codeWriter.writeBootload()

		for _, file := range files {
			path := filepath.Join(filePath, file.Name())

			// 拡張子チェック
			sPath := strings.Split(path, ".")
			if sPath[len(sPath)-1] != "vm" {
				log.Printf("%s is not vm file", file.Name())
				continue
			}
			f, err := os.Open(path)
			if err != nil {
				panic(err)
			}

			parser := newParser(f)
			defer parser.close()

			codeWriter.setFileName(f.Name())

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
		return
	}

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	codeWriter := newCodeWriter(f)
	defer codeWriter.close()

	parser := newParser(f)
	defer parser.close()

	codeWriter.setFileName(f.Name())

	for {
		if !parser.hasMoreCommands() {
			break
		}

		parser.advance()
		cm := parser.command

		if len(cm) > 0 {
			switch parser.commandType() {
			case C_ARITHMETIC:
				codeWriter.writeArthmethic(cm)
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
