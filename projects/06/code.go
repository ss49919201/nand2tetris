package main

// codeモジュール
type code struct{}

func newCode() *code {
	return &code{}
}

// destニーモニックをバイナリコードに変換する
func (c *code) dest(dm string) [3]byte {
	switch dm {
	case "null":
		return [3]byte{0, 0, 0}
	case "M":
		return [3]byte{0, 0, 1}
	case "D":
		return [3]byte{0, 1, 0}
	case "MD":
		return [3]byte{0, 1, 1}
	case "A":
		return [3]byte{1, 0, 0}
	case "AM":
		return [3]byte{1, 0, 1}
	case "AD":
		return [3]byte{1, 1, 0}
	case "AMD":
		return [3]byte{1, 1, 1}
	default:
		panic("not dest Mnemonic")
	}
}

// compニーモニックをバイナリコードに変換する
func (c *code) comp(cm string) [7]byte {
	switch cm {
	case "0":
		return [7]byte{0, 1, 0, 1, 0, 0, 0}
	case "1":
		return [7]byte{0, 1, 1, 1, 1, 1, 1}
	case "-1":
		return [7]byte{0, 1, 1, 1, 0, 1, 0}
	case "D":
		return [7]byte{0, 0, 0, 1, 1, 0, 0}
	case "A":
		return [7]byte{0, 1, 1, 0, 0, 0, 0}
	case "!D":
		return [7]byte{0, 0, 0, 1, 1, 0, 1}
	case "!A":
		return [7]byte{0, 1, 1, 0, 0, 0, 1}
	case "-D":
		return [7]byte{0, 0, 0, 1, 1, 1, 1}
	case "-A":
		return [7]byte{0, 1, 1, 0, 0, 1, 1}
	case "D+1":
		return [7]byte{0, 0, 1, 1, 1, 1, 1}
	case "A+1":
		return [7]byte{0, 1, 1, 0, 1, 1, 1}
	case "D-1":
		return [7]byte{0, 0, 0, 1, 1, 1, 0}
	case "A-1":
		return [7]byte{0, 1, 1, 0, 0, 1, 0}
	case "D+A":
		return [7]byte{0, 0, 0, 0, 0, 1, 0}
	case "D-A":
		return [7]byte{0, 0, 1, 0, 0, 1, 1}
	case "A-D":
		return [7]byte{0, 0, 0, 0, 1, 1, 1}
	case "D&A":
		return [7]byte{0, 0, 0, 0, 0, 0, 0}
	case "D|A":
		return [7]byte{0, 0, 1, 0, 1, 0, 1}
	case "M":
		return [7]byte{1, 1, 1, 0, 0, 0, 0}
	case "!M":
		return [7]byte{1, 1, 1, 0, 0, 0, 1}
	case "-M":
		return [7]byte{1, 1, 1, 0, 0, 1, 1}
	case "M+1":
		return [7]byte{1, 1, 1, 0, 1, 1, 1}
	case "M-1":
		return [7]byte{1, 1, 1, 0, 0, 1, 0}
	case "D+M":
		return [7]byte{1, 0, 0, 0, 0, 1, 0}
	case "D-M":
		return [7]byte{1, 0, 1, 0, 0, 1, 1}
	case "M-D":
		return [7]byte{1, 0, 0, 0, 1, 1, 1}
	case "D&M":
		return [7]byte{1, 0, 0, 0, 0, 0, 0}
	case "D|M":
		return [7]byte{1, 0, 1, 0, 1, 0, 1}
	default:
		panic("not jump Mnemonic")
	}
}

// jumpニーモニックをバイナリコードに変換する
func (c *code) jump(jm string) [3]byte {
	switch jm {
	case "null":
		return [3]byte{0, 0, 0}
	case "JGT":
		return [3]byte{0, 0, 1}
	case "JEQ":
		return [3]byte{0, 1, 0}
	case "JGE":
		return [3]byte{0, 1, 1}
	case "JLT":
		return [3]byte{1, 0, 0}
	case "JNE":
		return [3]byte{1, 0, 1}
	case "JLE":
		return [3]byte{1, 1, 0}
	case "JMP":
		return [3]byte{1, 1, 1}
	default:
		panic("not jump Mnemonic")
	}
}
