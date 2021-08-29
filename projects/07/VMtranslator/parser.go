package main

import (
	"bufio"
	"os"
	"strings"
)

type command int

const (
	C_ARITHMETIC command = iota
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IF
	C_FUNCTION
	C_RETURN
	C_CALL
)

type parser struct {
	input   *os.File
	scanner *bufio.Scanner
	command string
}

func newParser(file *os.File) *parser {
	scanner := bufio.NewScanner(file)
	return &parser{
		input:   file,
		scanner: scanner,
	}
}

// 入力にまだコマンドが存在するか
func (p *parser) hasMoreCommands() bool {
	return p.scanner.Scan()
}

// 次のコマンドを読み現在のコマンドにする
func (p *parser) advance() {
	// 左端の半角スペースを取り除く
	c := strings.Trim(p.scanner.Text(), " ")

	p.command = c
}

// 現在のコマンドの種類を返す
func (p *parser) commandType() command {
	c := p.command
	var cs []string
	if strings.Contains(c, " ") {
		cs = strings.Split(c, " ")
	} else {
		cs = append(cs, c)
	}

	switch string(cs[0]) {
	case "add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not":
		return C_ARITHMETIC
	case "push":
		return C_PUSH
	case "pop":
		return C_POP
	case "label":
		return C_LABEL
	case "goto":
		return C_GOTO
	case "if-goto":
		return C_IF
	case "function":
		return C_FUNCTION
	case "return":
		return C_RETURN
	case "call":
		return C_CALL
	default:
		return -1
	}
}

// コマンド自体を返す
func (p *parser) commandItself() string {
	c := trimComment(p.command)
	cs := splitByHalfSpace(c)
	return cs[0]
}

// 第1引数を返す
func (p *parser) arg1() string {
	if p.commandType() == C_RETURN {
		panic("invalid command C_RETURN")
	}

	c := trimComment(p.command)
	cs := splitByHalfSpace(c)

	if p.commandType() == C_ARITHMETIC {
		return cs[0]
	}
	return cs[1]
}

// 第2引数を返す
func (p *parser) arg2() string {
	switch p.commandType() {
	case C_POP, C_PUSH, C_FUNCTION, C_CALL:
		break
	default:
		panic("invalid command")
	}

	c := trimComment(p.command)
	cs := splitByHalfSpace(c)
	if len(cs) != 3 {
		panic("invalid command")
	}

	return cs[2]
}

// 入力ファイルをCloseする
func (p *parser) close() {
	p.input.Close()
}

func splitByHalfSpace(s string) []string {
	if strings.Contains(s, " ") {
		return strings.Split(s, " ")
	}

	return []string{s}
}

// コマンドと同じ行のコメントを取り除く
func trimComment(src string) string {
	ss := strings.Split(src, "//")
	ss[len(ss)-1] = strings.TrimRight(ss[len(ss)-1], " ")
	return ss[0]
}
