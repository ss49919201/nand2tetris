package main

import (
	"fmt"
	"strconv"
)

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func printBin(b []byte) {
	for i := 0; i < len(b); i++ {
		fmt.Print(b[i])
		if i == len(b)-1 {
			fmt.Println()
		}
	}
}

func makeSymbolTable(symbolTable *symbolTable) {
	parser := newParser()
	defer parser.file.Close()

	var commandIndex int
	for {
		cm := parser.command
		if len(cm) > 0 {
			switch parser.commandType() {
			case A_COMMAND, C_COMMAND:
				commandIndex++
			case L_COMMAND:
				symbolTable.addEntry(parser.symbol(), commandIndex)
			}
		}

		if parser.hasMoreCommands() {
			parser.advance()
			continue
		}
		break
	}
}

func makeBin(symbolTable *symbolTable) {
	parser := newParser()
	defer parser.file.Close()

	code := newCode()

	newSymbolIndex := 16

	for {
		cm := parser.command
		if len(cm) > 0 {
			switch parser.commandType() {
			case A_COMMAND:
				var s string
				// 数値の場合
				if i, err := strconv.Atoi(parser.symbol()); err == nil {
					// 15bitの2進数に変換
					s = fmt.Sprintf("%015b", i)
				} else {
					// シンボルの場合
					if !symbolTable.contains(parser.symbol()) {
						symbolTable.addEntry(parser.symbol(), newSymbolIndex)
						newSymbolIndex++
					}
					s = fmt.Sprintf("%015b", symbolTable.getAddress(parser.symbol()))
				}
				fmt.Println("0" + s)
			case C_COMMAND:
				b := make([]byte, 0, 16)
				b = append(b, []byte{1, 1, 1}...)

				for _, v := range code.comp(parser.comp()) {
					b = append(b, v)
				}

				for _, v := range code.dest(parser.dest()) {
					b = append(b, v)
				}

				for _, v := range code.jump(parser.jump()) {
					b = append(b, v)
				}

				printBin(b)
			case L_COMMAND:
			}
		}

		if parser.hasMoreCommands() {
			parser.advance()
			continue
		}
		break
	}
}

func main() {
	s := newSymbolTable()
	makeSymbolTable(s)
	makeBin(s)
}
