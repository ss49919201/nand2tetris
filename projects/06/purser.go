package main

import (
	"bufio"
	"os"
	"strings"
)

// purserモジュール
type purser struct {
	file    *os.File
	scanner *bufio.Scanner
	command string
}

type command int

const (
	A_COMMAND command = iota
	C_COMMAND
	L_COMMAND
)

// パーサーの初期化
func newPurser() *purser {
	// fileを開く
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	command := scanner.Text()
	return &purser{
		file:    file,
		scanner: scanner,
		command: command,
	}
}

// 入力にまだコマンドが存在するか
func (p *purser) hasMoreCommands() bool {
	return p.scanner.Scan()
}

// 次のコマンドを読み現在のコマンドにする
func (p *purser) advance() {
	p.command = p.scanner.Text()
}

// 現在のコマンドの種類を返す
func (p *purser) commandType() command {
	c := strings.Trim(p.command, " ")
	var cs []string
	if strings.Contains(c, " ") {
		cs = strings.Split(c, " ")
	} else {
		cs = append(cs, c)
	}

	switch s := string(cs[0][0]); s {
	case "@":
		return A_COMMAND
	case "A", "D", "M":
		return C_COMMAND
	case "(":
		return L_COMMAND
	default:
		if isInt(s) {
			return C_COMMAND
		}
		return -1
	}
}

// 現在のコマンド@Xxxまたは(Xxx)のXxxを文字列で返す
// コマンドがA_COMMANDまたはL_COMMANDの時だけ呼ぶ
func (p *purser) symbol() string {
	c := strings.Trim(p.command, " ")
	switch string(c[0]) {
	case "@":
		s := strings.TrimLeft(c, "@")
		return s
	case "(":
		s := strings.TrimLeft(c, "(")
		s = strings.TrimRight(s, ")")
		return s
	default:
		panic("not symbol")
	}
}

// 現在のC命令のdestニーモニックを返す（8つの可能性なので定数化する）
// コマンドがC_COMMANDの時だけ呼ぶ
func (p *purser) dest() string {
	c := strings.Trim(p.command, " ")
	cs := strings.Split(c, " ")

	if !strings.Contains(string(cs[0]), "=") {
		return "null"
	}
	s := strings.Split(string(cs[0]), "=")
	return s[0]
}

// 現在のC命令のcompニーモニックを返す（28個の可能性なので定数化する）
// コマンドがC_COMMANDの時だけ呼ぶ
func (p *purser) comp() string {
	c := strings.Trim(p.command, " ")
	cs := strings.Split(c, " ")

	if !strings.Contains(string(cs[0]), "=") {
		if !strings.Contains(string(cs[0]), ";") {
			return "null"
		}
		s := strings.Split(string(cs[0]), ";")
		return s[0]
	}
	s := strings.Split(string(cs[0]), "=")
	return s[1]
}

// 現在のC命令のjumpニーモニックを返す（8つの可能性なので定数化する）
// コマンドがC_COMMANDの時だけ呼ぶ
func (p *purser) jump() string {
	c := strings.Trim(p.command, " ")
	cs := strings.Split(c, " ")

	if !strings.Contains(string(cs[0]), ";") {
		return "null"
	}
	s := strings.Split(string(cs[0]), ";")
	return s[1]
}
