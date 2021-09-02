package main

import (
	"os"
	"strconv"
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

	codeWriter := newCodeWriter(file)
	defer codeWriter.close()

	// if isDir(file) {
	// 	files, err := file.ReadDir(0)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	for _, file := range files {
	// 		d := filepath.Join(filePath, file.Name())
	// 		if file.Type().IsDir() {
	// 			panic("can't translate directory")
	// 		}
	// 		f, err := os.Open(d)
	// 		if err != nil {
	// 			panic(err)
	// 		}

	// 		parser := newParser(f)
	// 		defer parser.close()
	// 	}
	// 	return
	// }

	parser := newParser(file)
	defer parser.close()

	codeWriter.setFileName(file.Name())
	codeWriter.close()

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
			}
		}
	}
}
