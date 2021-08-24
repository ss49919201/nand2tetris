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
	switch string(p.command[0]) {
	case "@":
		return A_COMMAND
	case "A", "D", "M":
		return C_COMMAND
	case "(":
		return L_COMMAND
	default:
		return -1
	}
}

// 現在のコマンド@Xxxまたは(Xxx)のXxxを文字列で返す
// コマンドがA_COMMANDまたはL_COMMANDの時だけ呼ぶ
func (p *purser) symbol() string {
	switch string(p.command[0]) {
	case "@":
		s := strings.TrimLeft(p.command, "@")
		return s
	case "(":
		s := strings.TrimLeft(p.command, "(")
		s = strings.TrimRight(p.command, ")")
		return s
	default:
		panic("not symbol")
	}
}

// 現在のC命令のdestニーモニックを返す（8つの可能性なので定数化する）
// コマンドがC_COMMANDの時だけ呼ぶ
func (p *purser) dest() string {
	if !strings.Contains(p.command, "=") {
		return "null"
	}
	s := strings.Split(p.command, "=")
	return s[0]
}

// 現在のC命令のcompニーモニックを返す（28個の可能性なので定数化する）
// コマンドがC_COMMANDの時だけ呼ぶ
func (p *purser) comp() string {
	if !strings.Contains(p.command, "=") {
		return "null"
	}
	s := strings.Split(p.command, "=")
	return s[1]
}

// 現在のC命令のjumpニーモニックを返す（8つの可能性なので定数化する）
// コマンドがC_COMMANDの時だけ呼ぶ
func (p *purser) jump() string {
	if !strings.Contains(p.command, ";") {
		return "null"
	}
	s := strings.Split(p.command, ";")
	return s[1]
}
