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
	purser := newPurser()
	defer purser.file.Close()

	var commandIndex int
	for {
		cm := purser.command
		if len(cm) > 0 {
			switch purser.commandType() {
			case A_COMMAND, C_COMMAND:
				commandIndex++
			case L_COMMAND:
				symbolTable.addEntry(purser.symbol(), commandIndex)
			}
		}

		if purser.hasMoreCommands() {
			purser.advance()
			continue
		}
		break
	}
}

func makeBin(symbolTable *symbolTable) {
	purser := newPurser()
	defer purser.file.Close()

	code := newCode()

	newSymbolIndex := 16

	for {
		cm := purser.command
		if len(cm) > 0 {
			switch purser.commandType() {
			case A_COMMAND:
				var s string
				// 数値の場合
				if i, err := strconv.Atoi(purser.symbol()); err == nil {
					// 15bitの2進数に変換
					s = fmt.Sprintf("%015b", i)
				} else {
					// シンボルの場合
					if !symbolTable.contains(purser.symbol()) {
						symbolTable.addEntry(purser.symbol(), newSymbolIndex)
						newSymbolIndex++
					}
					s = fmt.Sprintf("%015b", symbolTable.getAddress(purser.symbol()))
				}
				fmt.Println("0" + s)
			case C_COMMAND:
				b := make([]byte, 0, 16)
				b = append(b, []byte{1, 1, 1}...)

				for _, v := range code.comp(purser.comp()) {
					b = append(b, v)
				}

				for _, v := range code.dest(purser.dest()) {
					b = append(b, v)
				}

				for _, v := range code.jump(purser.jump()) {
					b = append(b, v)
				}

				printBin(b)
			case L_COMMAND:
			}
		}

		if purser.hasMoreCommands() {
			purser.advance()
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
