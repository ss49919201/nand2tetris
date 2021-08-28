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
	file    *os.File
	scanner *bufio.Scanner
	command string
}

func newParser() *parser {
	// fileを開く
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	return &parser{
		file:    file,
		scanner: scanner,
	}
}

// 入力にまだコマンドが存在するか
func (p *parser) hasMoreCommands() bool {
	return p.scanner.Scan()
}

// 次のコマンドを読み現在のコマンドにする
func (p *parser) advance() {
	p.command = strings.Trim(p.scanner.Text(), " ")
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
		panic("invalid command")
	}
}

func (p *parser) arg1() string {
	if p.commandType() == C_RETURN {
		panic("invalid command C_RETURN")
	}

	c := p.command
	var cs []string
	if strings.Contains(c, " ") {
		cs = strings.Split(c, " ")
	} else {
		cs = append(cs, c)
	}

	if p.commandType() == C_ARITHMETIC {
		return cs[0]
	}
	return cs[1]
}

func (p *parser) arg2() string {
	switch p.commandType() {
	case C_POP, C_PUSH, C_FUNCTION, C_CALL:
		break
	default:
		panic("invalid command")
	}

	c := p.command
	var cs []string
	if strings.Contains(c, " ") {
		cs = strings.Split(c, " ")
	} else {
		cs = append(cs, c)
	}

	return cs[2]
}

// 入力にまだコマンドが存在するか
func (p *parser) close() {
	p.file.Close()
}
