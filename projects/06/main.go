package main

import (
	"fmt"
	"strconv"
)

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func main() {
	purser := newPurser()
	defer purser.file.Close()

	code := newCode()

	println("perse command...")
	for {
		cm := purser.command
		if len(cm) > 0 {
			switch purser.commandType() {
			case A_COMMAND:
				println()
				// 数値の場合
				if i, err := strconv.Atoi(purser.symbol()); err == nil {
					// 15bitの2進数に変換
					s := fmt.Sprintf("%015b", i)
					println("value", s)
					println()
				} else {
					// シンボルの場合
					fmt.Println("symbol=", purser.symbol())
				}
			case C_COMMAND:
				println()
				println(cm)
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

				fmt.Println(b)
			case L_COMMAND:
			}
		}

		if purser.hasMoreCommands() {
			purser.advance()
			continue
		}
		println("break loop")
		break
	}
}
